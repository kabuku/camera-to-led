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

// NewGetPrinterBedPreHeatParams creates a new GetPrinterBedPreHeatParams object
// with the default values initialized.
func NewGetPrinterBedPreHeatParams() *GetPrinterBedPreHeatParams {

	return &GetPrinterBedPreHeatParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterBedPreHeatParamsWithTimeout creates a new GetPrinterBedPreHeatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterBedPreHeatParamsWithTimeout(timeout time.Duration) *GetPrinterBedPreHeatParams {

	return &GetPrinterBedPreHeatParams{

		timeout: timeout,
	}
}

// NewGetPrinterBedPreHeatParamsWithContext creates a new GetPrinterBedPreHeatParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterBedPreHeatParamsWithContext(ctx context.Context) *GetPrinterBedPreHeatParams {

	return &GetPrinterBedPreHeatParams{

		Context: ctx,
	}
}

// NewGetPrinterBedPreHeatParamsWithHTTPClient creates a new GetPrinterBedPreHeatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterBedPreHeatParamsWithHTTPClient(client *http.Client) *GetPrinterBedPreHeatParams {

	return &GetPrinterBedPreHeatParams{
		HTTPClient: client,
	}
}

/*GetPrinterBedPreHeatParams contains all the parameters to send to the API endpoint
for the get printer bed pre heat operation typically these are written to a http.Request
*/
type GetPrinterBedPreHeatParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer bed pre heat params
func (o *GetPrinterBedPreHeatParams) WithTimeout(timeout time.Duration) *GetPrinterBedPreHeatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer bed pre heat params
func (o *GetPrinterBedPreHeatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer bed pre heat params
func (o *GetPrinterBedPreHeatParams) WithContext(ctx context.Context) *GetPrinterBedPreHeatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer bed pre heat params
func (o *GetPrinterBedPreHeatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer bed pre heat params
func (o *GetPrinterBedPreHeatParams) WithHTTPClient(client *http.Client) *GetPrinterBedPreHeatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer bed pre heat params
func (o *GetPrinterBedPreHeatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterBedPreHeatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}