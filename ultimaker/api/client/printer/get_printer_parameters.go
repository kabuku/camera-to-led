package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrinterParams creates a new GetPrinterParams object
// with the default values initialized.
func NewGetPrinterParams() *GetPrinterParams {

	return &GetPrinterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterParamsWithTimeout creates a new GetPrinterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterParamsWithTimeout(timeout time.Duration) *GetPrinterParams {

	return &GetPrinterParams{

		timeout: timeout,
	}
}

// NewGetPrinterParamsWithContext creates a new GetPrinterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterParamsWithContext(ctx context.Context) *GetPrinterParams {

	return &GetPrinterParams{

		Context: ctx,
	}
}

// NewGetPrinterParamsWithHTTPClient creates a new GetPrinterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterParamsWithHTTPClient(client *http.Client) *GetPrinterParams {

	return &GetPrinterParams{
		HTTPClient: client,
	}
}

/*GetPrinterParams contains all the parameters to send to the API endpoint
for the get printer operation typically these are written to a http.Request
*/
type GetPrinterParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer params
func (o *GetPrinterParams) WithTimeout(timeout time.Duration) *GetPrinterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer params
func (o *GetPrinterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer params
func (o *GetPrinterParams) WithContext(ctx context.Context) *GetPrinterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer params
func (o *GetPrinterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer params
func (o *GetPrinterParams) WithHTTPClient(client *http.Client) *GetPrinterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer params
func (o *GetPrinterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}