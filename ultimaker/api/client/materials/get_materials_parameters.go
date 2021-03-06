package materials

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

// NewGetMaterialsParams creates a new GetMaterialsParams object
// with the default values initialized.
func NewGetMaterialsParams() *GetMaterialsParams {

	return &GetMaterialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaterialsParamsWithTimeout creates a new GetMaterialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMaterialsParamsWithTimeout(timeout time.Duration) *GetMaterialsParams {

	return &GetMaterialsParams{

		timeout: timeout,
	}
}

// NewGetMaterialsParamsWithContext creates a new GetMaterialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMaterialsParamsWithContext(ctx context.Context) *GetMaterialsParams {

	return &GetMaterialsParams{

		Context: ctx,
	}
}

// NewGetMaterialsParamsWithHTTPClient creates a new GetMaterialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMaterialsParamsWithHTTPClient(client *http.Client) *GetMaterialsParams {

	return &GetMaterialsParams{
		HTTPClient: client,
	}
}

/*GetMaterialsParams contains all the parameters to send to the API endpoint
for the get materials operation typically these are written to a http.Request
*/
type GetMaterialsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get materials params
func (o *GetMaterialsParams) WithTimeout(timeout time.Duration) *GetMaterialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get materials params
func (o *GetMaterialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get materials params
func (o *GetMaterialsParams) WithContext(ctx context.Context) *GetMaterialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get materials params
func (o *GetMaterialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get materials params
func (o *GetMaterialsParams) WithHTTPClient(client *http.Client) *GetMaterialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get materials params
func (o *GetMaterialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaterialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
