package system

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

// NewGetSystemVariantParams creates a new GetSystemVariantParams object
// with the default values initialized.
func NewGetSystemVariantParams() *GetSystemVariantParams {

	return &GetSystemVariantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemVariantParamsWithTimeout creates a new GetSystemVariantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemVariantParamsWithTimeout(timeout time.Duration) *GetSystemVariantParams {

	return &GetSystemVariantParams{

		timeout: timeout,
	}
}

// NewGetSystemVariantParamsWithContext creates a new GetSystemVariantParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemVariantParamsWithContext(ctx context.Context) *GetSystemVariantParams {

	return &GetSystemVariantParams{

		Context: ctx,
	}
}

// NewGetSystemVariantParamsWithHTTPClient creates a new GetSystemVariantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemVariantParamsWithHTTPClient(client *http.Client) *GetSystemVariantParams {

	return &GetSystemVariantParams{
		HTTPClient: client,
	}
}

/*GetSystemVariantParams contains all the parameters to send to the API endpoint
for the get system variant operation typically these are written to a http.Request
*/
type GetSystemVariantParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system variant params
func (o *GetSystemVariantParams) WithTimeout(timeout time.Duration) *GetSystemVariantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system variant params
func (o *GetSystemVariantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system variant params
func (o *GetSystemVariantParams) WithContext(ctx context.Context) *GetSystemVariantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system variant params
func (o *GetSystemVariantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system variant params
func (o *GetSystemVariantParams) WithHTTPClient(client *http.Client) *GetSystemVariantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system variant params
func (o *GetSystemVariantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemVariantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
