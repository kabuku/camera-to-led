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

// NewGetSystemTypeParams creates a new GetSystemTypeParams object
// with the default values initialized.
func NewGetSystemTypeParams() *GetSystemTypeParams {

	return &GetSystemTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemTypeParamsWithTimeout creates a new GetSystemTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemTypeParamsWithTimeout(timeout time.Duration) *GetSystemTypeParams {

	return &GetSystemTypeParams{

		timeout: timeout,
	}
}

// NewGetSystemTypeParamsWithContext creates a new GetSystemTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemTypeParamsWithContext(ctx context.Context) *GetSystemTypeParams {

	return &GetSystemTypeParams{

		Context: ctx,
	}
}

// NewGetSystemTypeParamsWithHTTPClient creates a new GetSystemTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemTypeParamsWithHTTPClient(client *http.Client) *GetSystemTypeParams {

	return &GetSystemTypeParams{
		HTTPClient: client,
	}
}

/*GetSystemTypeParams contains all the parameters to send to the API endpoint
for the get system type operation typically these are written to a http.Request
*/
type GetSystemTypeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system type params
func (o *GetSystemTypeParams) WithTimeout(timeout time.Duration) *GetSystemTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system type params
func (o *GetSystemTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system type params
func (o *GetSystemTypeParams) WithContext(ctx context.Context) *GetSystemTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system type params
func (o *GetSystemTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system type params
func (o *GetSystemTypeParams) WithHTTPClient(client *http.Client) *GetSystemTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system type params
func (o *GetSystemTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
