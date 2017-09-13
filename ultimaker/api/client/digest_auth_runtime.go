package client

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"

	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"path/filepath"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DefaultTimeout the default request timeout
var DefaultTimeout = 30 * time.Second

// Runtime represents an API client that uses the transport
// to make http requests based on a swagger specification.
type Runtime struct {
	DefaultMediaType      string
	DefaultAuthentication runtime.ClientAuthInfoWriter
	Consumers             map[string]runtime.Consumer
	Producers             map[string]runtime.Producer

	Transport http.RoundTripper
	Jar       http.CookieJar
	//Spec      *spec.Document
	Host     string
	BasePath string
	Formats  strfmt.Registry
	Debug    bool
	Context  context.Context

	clientOnce *sync.Once
	client     *http.Client
	schemes    []string
	do         func(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error)
	username   string
	password   string
}

// New creates a new default runtime for a swagger api runtime.Client
func NewDigestAuthRuntime(host, basePath string, schemes []string, username, password string) *Runtime {
	var rt Runtime
	rt.DefaultMediaType = runtime.JSONMime

	// TODO: actually infer this stuff from the spec
	rt.Consumers = map[string]runtime.Consumer{
		runtime.JSONMime:    runtime.JSONConsumer(),
		runtime.XMLMime:     runtime.XMLConsumer(),
		runtime.TextMime:    runtime.TextConsumer(),
		runtime.DefaultMime: runtime.ByteStreamConsumer(),
	}
	rt.Producers = map[string]runtime.Producer{
		runtime.JSONMime:    runtime.JSONProducer(),
		runtime.XMLMime:     runtime.XMLProducer(),
		runtime.TextMime:    runtime.TextProducer(),
		runtime.DefaultMime: runtime.ByteStreamProducer(),
	}
	rt.Transport = http.DefaultTransport
	rt.Jar = nil
	rt.Host = host
	rt.BasePath = basePath
	rt.Context = context.Background()
	rt.clientOnce = new(sync.Once)
	if !strings.HasPrefix(rt.BasePath, "/") {
		rt.BasePath = "/" + rt.BasePath
	}
	rt.Debug = len(os.Getenv("DEBUG")) > 0
	if len(schemes) > 0 {
		rt.schemes = schemes
	}
	rt.do = ctxhttp.Do
	rt.username = username
	rt.password = password
	return &rt
}

// NewWithClient allows you to create a new transport with a configured http.Client
func NewWithClient(host, basePath string, schemes []string, client *http.Client, username, password string) *Runtime {
	rt := NewDigestAuthRuntime(host, basePath, schemes, username, password)
	if client != nil {
		rt.clientOnce.Do(func() {
			rt.client = client
		})
	}
	return rt
}

func (r *Runtime) pickScheme(schemes []string) string {
	if v := r.selectScheme(r.schemes); v != "" {
		return v
	}
	if v := r.selectScheme(schemes); v != "" {
		return v
	}
	return "http"
}

func (r *Runtime) selectScheme(schemes []string) string {
	schLen := len(schemes)
	if schLen == 0 {
		return ""
	}

	scheme := schemes[0]
	// prefer https, but skip when not possible
	if scheme != "https" && schLen > 1 {
		for _, sch := range schemes {
			if sch == "https" {
				scheme = sch
				break
			}
		}
	}
	return scheme
}

// Submit a request and when there is a body on success it will turn that into the result
// all other things are turned into an api error for swagger which retains the status code
func (r *Runtime) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	params, readResponse, auth := operation.Params, operation.Reader, operation.AuthInfo

	request, err := newRequest(operation.Method, operation.PathPattern, params)
	if err != nil {
		return nil, err
	}

	var accept []string
	for _, mimeType := range operation.ProducesMediaTypes {
		accept = append(accept, mimeType)
	}
	request.SetHeaderParam(runtime.HeaderAccept, accept...)

	if auth == nil && r.DefaultAuthentication != nil {
		auth = r.DefaultAuthentication
	}
	if auth != nil {
		if err := auth.AuthenticateRequest(request, r.Formats); err != nil {
			return nil, err
		}
	}

	// TODO: pick appropriate media type
	cmt := r.DefaultMediaType
	for _, mediaType := range operation.ConsumesMediaTypes {
		// Pick first non-empty media type
		if mediaType != "" {
			cmt = mediaType
			break
		}
	}
	req, res, err := r.doRequest(operation, request, cmt) // make requests, by default follows 10 redirects before failing
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 403 || res.StatusCode == 401 {
		authHeader := CreateDigestAuthHeader(response{res}, req.Method, req.URL.Path, r.username, r.password)
		request.SetHeaderParam("Authorization", authHeader)
		req, res, err = r.doRequest(operation, request, cmt) // make requests, by default follows 10 redirects before failing

		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
	}

	ct := res.Header.Get(runtime.HeaderContentType)
	if ct == "" { // this should really really never occur
		ct = r.DefaultMediaType
	}

	mt, _, err := mime.ParseMediaType(ct)
	if err != nil {
		return nil, fmt.Errorf("parse content type: %s", err)
	}

	cons, ok := r.Consumers[mt]
	if !ok {
		// scream about not knowing what to do
		return nil, fmt.Errorf("no consumer: %q", ct)
	}

	return readResponse.ReadResponse(response{res}, cons)
}

func (r *Runtime) doRequest(operation *runtime.ClientOperation, request *request, cmt string) (*http.Request, *http.Response, error) {
	req, err := request.BuildHTTP(cmt, r.Producers, r.Formats)
	if err != nil {
		return nil, nil, err
	}
	req.URL.Scheme = r.pickScheme(operation.Schemes)
	req.URL.Host = r.Host
	var reinstateSlash bool
	if req.URL.Path != "" && req.URL.Path != "/" && req.URL.Path[len(req.URL.Path)-1] == '/' {
		reinstateSlash = true
	}
	req.URL.Path = path.Join(r.BasePath, req.URL.Path)
	if reinstateSlash {
		req.URL.Path = req.URL.Path + "/"
	}

	r.clientOnce.Do(func() {
		r.client = &http.Client{
			Transport: r.Transport,
			Jar:       r.Jar,
		}
	})

	if r.Debug {
		b, err2 := httputil.DumpRequestOut(req, true)
		if err2 != nil {
			return nil, nil, err2
		}
		fmt.Fprintln(os.Stderr, string(b))
	}

	var hasTimeout bool
	pctx := operation.Context
	if pctx == nil {
		pctx = r.Context
	} else {
		hasTimeout = true
	}
	if pctx == nil {
		pctx = context.Background()
	}
	var ctx context.Context
	var cancel context.CancelFunc
	if hasTimeout {
		ctx, cancel = context.WithCancel(pctx)
	} else {
		ctx, cancel = context.WithTimeout(pctx, request.timeout)
	}
	defer cancel()

	client := operation.Client
	if client == nil {
		client = r.client
	}
	if r.do == nil {
		r.do = ctxhttp.Do
	}

	res, err := r.do(ctx, client, req) // make requests, by default follows 10 redirects before failing
	if err != nil {
		return nil, nil, err
	}
	if r.Debug {
		b, err2 := httputil.DumpResponse(res, true)
		if err2 != nil {
			return nil, nil, err2
		}
		fmt.Fprintln(os.Stderr, string(b))
	}

	return req, res, err
}

// NewRequest creates a new swagger http client request
func newRequest(method, pathPattern string, writer runtime.ClientRequestWriter) (*request, error) {
	return &request{
		pathPattern: pathPattern,
		method:      method,
		writer:      writer,
		header:      make(http.Header),
		query:       make(url.Values),
		timeout:     DefaultTimeout,
	}, nil
}

// Request represents a swagger client request.
//
// This Request struct converts to a HTTP request.
// There might be others that convert to other transports.
// There is no error checking here, it is assumed to be used after a spec has been validated.
// so impossible combinations should not arise (hopefully).
//
// The main purpose of this struct is to hide the machinery of adding params to a transport request.
// The generated code only implements what is necessary to turn a param into a valid value for these methods.
type request struct {
	pathPattern string
	method      string
	writer      runtime.ClientRequestWriter

	pathParams map[string]string
	header     http.Header
	query      url.Values
	formFields url.Values
	fileFields map[string]runtime.NamedReadCloser
	payload    interface{}
	timeout    time.Duration
}

// BuildHTTP creates a new http request based on the data from the params
func (r *request) BuildHTTP(mediaType string, producers map[string]runtime.Producer, registry strfmt.Registry) (*http.Request, error) {
	// build the data
	if err := r.writer.WriteToRequest(r, registry); err != nil {
		return nil, err
	}

	// create http request
	path := r.pathPattern
	for k, v := range r.pathParams {
		path = strings.Replace(path, "{"+k+"}", v, -1)
	}

	var body io.ReadCloser
	var pr *io.PipeReader
	var pw *io.PipeWriter
	buf := bytes.NewBuffer(nil)
	if r.payload != nil || len(r.formFields) > 0 || len(r.fileFields) > 0 {
		body = ioutil.NopCloser(buf)
		if r.fileFields != nil {
			pr, pw = io.Pipe()
			body = pr
		}
	}
	req, err := http.NewRequest(r.method, path, body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = r.query.Encode()
	req.Header = r.header

	// check if this is a form type request
	if len(r.formFields) > 0 || len(r.fileFields) > 0 {
		// check if this is multipart
		if len(r.fileFields) > 0 {
			mp := multipart.NewWriter(pw)
			req.Header.Set(runtime.HeaderContentType, mp.FormDataContentType())

			go func() {
				defer func() {
					mp.Close()
					pw.Close()
				}()

				for fn, v := range r.formFields {
					if len(v) > 0 {
						if err := mp.WriteField(fn, v[0]); err != nil {
							pw.CloseWithError(err)
							log.Println(err)
						}
					}
				}

				for fn, f := range r.fileFields {
					wrtr, err := mp.CreateFormFile(fn, filepath.Base(f.Name()))
					if err != nil {
						pw.CloseWithError(err)
						log.Println(err)
					}
					defer func() {
						for _, ff := range r.fileFields {
							ff.Close()
						}

					}()
					if _, err := io.Copy(wrtr, f); err != nil {
						pw.CloseWithError(err)
						log.Println(err)
					}
				}

			}()
			return req, nil
		}

		req.Header.Set(runtime.HeaderContentType, mediaType)
		formString := r.formFields.Encode()
		// set content length before writing to the buffer
		req.ContentLength = int64(len(formString))
		// write the form values as the body
		buf.WriteString(formString)
		return req, nil
	}

	// if there is payload, use the producer to write the payload, and then
	// set the header to the content-type appropriate for the payload produced
	if r.payload != nil {
		// TODO: infer most appropriate content type based on the producer used,
		// and the `consumers` section of the spec/operation
		req.Header.Set(runtime.HeaderContentType, mediaType)
		if rdr, ok := r.payload.(io.ReadCloser); ok {
			req.Body = rdr
			return req, nil
		}

		if rdr, ok := r.payload.(io.Reader); ok {
			req.Body = ioutil.NopCloser(rdr)
			return req, nil
		}

		// set the content length of the request or else a chunked transfer is
		// declared, and this corrupts outgoing JSON payloads. the content's
		// length must be set prior to the body being written per the spec at
		// https://golang.org/pkg/net/http
		//
		//     If Body is present, Content-Length is <= 0 and TransferEncoding
		//     hasn't been set to "identity", Write adds
		//     "Transfer-Encoding: chunked" to the header. Body is closed
		//     after it is sent.
		//
		// to that end a temporary buffer, b, is created to produce the payload
		// body, and then its size is used to set the request's content length
		var b bytes.Buffer
		producer := producers[mediaType]
		if err := producer.Produce(&b, r.payload); err != nil {
			return nil, err
		}
		req.ContentLength = int64(b.Len())
		if _, err := buf.Write(b.Bytes()); err != nil {
			return nil, err
		}
	}

	if runtime.CanHaveBody(req.Method) && req.Body == nil && req.Header.Get(runtime.HeaderContentType) == "" {
		req.Header.Set(runtime.HeaderContentType, mediaType)
	}

	return req, nil
}

// SetHeaderParam adds a header param to the request
// when there is only 1 value provided for the varargs, it will set it.
// when there are several values provided for the varargs it will add it (no overriding)
func (r *request) SetHeaderParam(name string, values ...string) error {
	if r.header == nil {
		r.header = make(http.Header)
	}
	r.header[http.CanonicalHeaderKey(name)] = values
	return nil
}

// SetQueryParam adds a query param to the request
// when there is only 1 value provided for the varargs, it will set it.
// when there are several values provided for the varargs it will add it (no overriding)
func (r *request) SetQueryParam(name string, values ...string) error {
	if r.query == nil {
		r.query = make(url.Values)
	}
	r.query[name] = values
	return nil
}

// SetFormParam adds a forn param to the request
// when there is only 1 value provided for the varargs, it will set it.
// when there are several values provided for the varargs it will add it (no overriding)
func (r *request) SetFormParam(name string, values ...string) error {
	if r.formFields == nil {
		r.formFields = make(url.Values)
	}
	r.formFields[name] = values
	return nil
}

// SetPathParam adds a path param to the request
func (r *request) SetPathParam(name string, value string) error {
	if r.pathParams == nil {
		r.pathParams = make(map[string]string)
	}

	r.pathParams[name] = value
	return nil
}

// SetFileParam adds a file param to the request
func (r *request) SetFileParam(name string, file runtime.NamedReadCloser) error {
	if actualFile, ok := file.(*os.File); ok {
		fi, err := os.Stat(actualFile.Name())
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return fmt.Errorf("%q is a directory, only files are supported", file.Name())
		}
	}

	if r.fileFields == nil {
		r.fileFields = make(map[string]runtime.NamedReadCloser)
	}
	if r.formFields == nil {
		r.formFields = make(url.Values)
	}

	r.fileFields[name] = file
	return nil
}

// SetBodyParam sets a body parameter on the request.
// This does not yet serialze the object, this happens as late as possible.
func (r *request) SetBodyParam(payload interface{}) error {
	r.payload = payload
	return nil
}

// SetTimeout sets the timeout for a request
func (r *request) SetTimeout(timeout time.Duration) error {
	r.timeout = timeout
	return nil
}

var _ runtime.ClientResponse = response{}

type response struct {
	resp *http.Response
}

func (r response) Code() int {
	return r.resp.StatusCode
}

func (r response) Message() string {
	return r.resp.Status
}

func (r response) GetHeader(name string) string {
	return r.resp.Header.Get(name)
}

func (r response) Body() io.ReadCloser {
	return r.resp.Body
}
