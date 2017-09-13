package print_job

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

// NewGetPrintJobTimeTotalParams creates a new GetPrintJobTimeTotalParams object
// with the default values initialized.
func NewGetPrintJobTimeTotalParams() *GetPrintJobTimeTotalParams {

	return &GetPrintJobTimeTotalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrintJobTimeTotalParamsWithTimeout creates a new GetPrintJobTimeTotalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrintJobTimeTotalParamsWithTimeout(timeout time.Duration) *GetPrintJobTimeTotalParams {

	return &GetPrintJobTimeTotalParams{

		timeout: timeout,
	}
}

// NewGetPrintJobTimeTotalParamsWithContext creates a new GetPrintJobTimeTotalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrintJobTimeTotalParamsWithContext(ctx context.Context) *GetPrintJobTimeTotalParams {

	return &GetPrintJobTimeTotalParams{

		Context: ctx,
	}
}

// NewGetPrintJobTimeTotalParamsWithHTTPClient creates a new GetPrintJobTimeTotalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrintJobTimeTotalParamsWithHTTPClient(client *http.Client) *GetPrintJobTimeTotalParams {

	return &GetPrintJobTimeTotalParams{
		HTTPClient: client,
	}
}

/*GetPrintJobTimeTotalParams contains all the parameters to send to the API endpoint
for the get print job time total operation typically these are written to a http.Request
*/
type GetPrintJobTimeTotalParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get print job time total params
func (o *GetPrintJobTimeTotalParams) WithTimeout(timeout time.Duration) *GetPrintJobTimeTotalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get print job time total params
func (o *GetPrintJobTimeTotalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get print job time total params
func (o *GetPrintJobTimeTotalParams) WithContext(ctx context.Context) *GetPrintJobTimeTotalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get print job time total params
func (o *GetPrintJobTimeTotalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get print job time total params
func (o *GetPrintJobTimeTotalParams) WithHTTPClient(client *http.Client) *GetPrintJobTimeTotalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get print job time total params
func (o *GetPrintJobTimeTotalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrintJobTimeTotalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
