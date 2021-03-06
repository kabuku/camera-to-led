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

// NewGetPrinterLedParams creates a new GetPrinterLedParams object
// with the default values initialized.
func NewGetPrinterLedParams() *GetPrinterLedParams {

	return &GetPrinterLedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterLedParamsWithTimeout creates a new GetPrinterLedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterLedParamsWithTimeout(timeout time.Duration) *GetPrinterLedParams {

	return &GetPrinterLedParams{

		timeout: timeout,
	}
}

// NewGetPrinterLedParamsWithContext creates a new GetPrinterLedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterLedParamsWithContext(ctx context.Context) *GetPrinterLedParams {

	return &GetPrinterLedParams{

		Context: ctx,
	}
}

// NewGetPrinterLedParamsWithHTTPClient creates a new GetPrinterLedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterLedParamsWithHTTPClient(client *http.Client) *GetPrinterLedParams {

	return &GetPrinterLedParams{
		HTTPClient: client,
	}
}

/*GetPrinterLedParams contains all the parameters to send to the API endpoint
for the get printer led operation typically these are written to a http.Request
*/
type GetPrinterLedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer led params
func (o *GetPrinterLedParams) WithTimeout(timeout time.Duration) *GetPrinterLedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer led params
func (o *GetPrinterLedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer led params
func (o *GetPrinterLedParams) WithContext(ctx context.Context) *GetPrinterLedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer led params
func (o *GetPrinterLedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer led params
func (o *GetPrinterLedParams) WithHTTPClient(client *http.Client) *GetPrinterLedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer led params
func (o *GetPrinterLedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterLedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
