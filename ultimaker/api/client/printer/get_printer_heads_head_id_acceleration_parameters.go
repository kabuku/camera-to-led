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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrinterHeadsHeadIDAccelerationParams creates a new GetPrinterHeadsHeadIDAccelerationParams object
// with the default values initialized.
func NewGetPrinterHeadsHeadIDAccelerationParams() *GetPrinterHeadsHeadIDAccelerationParams {
	var ()
	return &GetPrinterHeadsHeadIDAccelerationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterHeadsHeadIDAccelerationParamsWithTimeout creates a new GetPrinterHeadsHeadIDAccelerationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterHeadsHeadIDAccelerationParamsWithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDAccelerationParams {
	var ()
	return &GetPrinterHeadsHeadIDAccelerationParams{

		timeout: timeout,
	}
}

// NewGetPrinterHeadsHeadIDAccelerationParamsWithContext creates a new GetPrinterHeadsHeadIDAccelerationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterHeadsHeadIDAccelerationParamsWithContext(ctx context.Context) *GetPrinterHeadsHeadIDAccelerationParams {
	var ()
	return &GetPrinterHeadsHeadIDAccelerationParams{

		Context: ctx,
	}
}

// NewGetPrinterHeadsHeadIDAccelerationParamsWithHTTPClient creates a new GetPrinterHeadsHeadIDAccelerationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterHeadsHeadIDAccelerationParamsWithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDAccelerationParams {
	var ()
	return &GetPrinterHeadsHeadIDAccelerationParams{
		HTTPClient: client,
	}
}

/*GetPrinterHeadsHeadIDAccelerationParams contains all the parameters to send to the API endpoint
for the get printer heads head ID acceleration operation typically these are written to a http.Request
*/
type GetPrinterHeadsHeadIDAccelerationParams struct {

	/*HeadID
	  ID of head of which to get the default acceleration of. Note that this speed also has a Z component. This API assumes that the head is the only part that moves.

	*/
	HeadID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) WithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDAccelerationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) WithContext(ctx context.Context) *GetPrinterHeadsHeadIDAccelerationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) WithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDAccelerationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeadID adds the headID to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) WithHeadID(headID int64) *GetPrinterHeadsHeadIDAccelerationParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the get printer heads head ID acceleration params
func (o *GetPrinterHeadsHeadIDAccelerationParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterHeadsHeadIDAccelerationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param head_id
	if err := r.SetPathParam("head_id", swag.FormatInt64(o.HeadID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
