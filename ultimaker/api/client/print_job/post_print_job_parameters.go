package print_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostPrintJobParams creates a new PostPrintJobParams object
// with the default values initialized.
func NewPostPrintJobParams() *PostPrintJobParams {
	var ()
	return &PostPrintJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrintJobParamsWithTimeout creates a new PostPrintJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrintJobParamsWithTimeout(timeout time.Duration) *PostPrintJobParams {
	var ()
	return &PostPrintJobParams{

		timeout: timeout,
	}
}

// NewPostPrintJobParamsWithContext creates a new PostPrintJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrintJobParamsWithContext(ctx context.Context) *PostPrintJobParams {
	var ()
	return &PostPrintJobParams{

		Context: ctx,
	}
}

// NewPostPrintJobParamsWithHTTPClient creates a new PostPrintJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrintJobParamsWithHTTPClient(client *http.Client) *PostPrintJobParams {
	var ()
	return &PostPrintJobParams{
		HTTPClient: client,
	}
}

/*PostPrintJobParams contains all the parameters to send to the API endpoint
for the post print job operation typically these are written to a http.Request
*/
type PostPrintJobParams struct {

	/*File
	  File that needs to be printed (.gcode, .gcode.gz)

	*/
	File os.File
	/*Jobname
	  Name of the print job.

	*/
	Jobname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post print job params
func (o *PostPrintJobParams) WithTimeout(timeout time.Duration) *PostPrintJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post print job params
func (o *PostPrintJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post print job params
func (o *PostPrintJobParams) WithContext(ctx context.Context) *PostPrintJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post print job params
func (o *PostPrintJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post print job params
func (o *PostPrintJobParams) WithHTTPClient(client *http.Client) *PostPrintJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post print job params
func (o *PostPrintJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the post print job params
func (o *PostPrintJobParams) WithFile(file os.File) *PostPrintJobParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the post print job params
func (o *PostPrintJobParams) SetFile(file os.File) {
	o.File = file
}

// WithJobname adds the jobname to the post print job params
func (o *PostPrintJobParams) WithJobname(jobname string) *PostPrintJobParams {
	o.SetJobname(jobname)
	return o
}

// SetJobname adds the jobname to the post print job params
func (o *PostPrintJobParams) SetJobname(jobname string) {
	o.Jobname = jobname
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrintJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param file
	if err := r.SetFileParam("file", &o.File); err != nil {
		return err
	}

	// form param jobname
	frJobname := o.Jobname
	fJobname := frJobname
	if fJobname != "" {
		if err := r.SetFormParam("jobname", fJobname); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}