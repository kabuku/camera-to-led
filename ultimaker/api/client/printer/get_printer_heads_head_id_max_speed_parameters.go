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

// NewGetPrinterHeadsHeadIDMaxSpeedParams creates a new GetPrinterHeadsHeadIDMaxSpeedParams object
// with the default values initialized.
func NewGetPrinterHeadsHeadIDMaxSpeedParams() *GetPrinterHeadsHeadIDMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDMaxSpeedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterHeadsHeadIDMaxSpeedParamsWithTimeout creates a new GetPrinterHeadsHeadIDMaxSpeedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterHeadsHeadIDMaxSpeedParamsWithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDMaxSpeedParams{

		timeout: timeout,
	}
}

// NewGetPrinterHeadsHeadIDMaxSpeedParamsWithContext creates a new GetPrinterHeadsHeadIDMaxSpeedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterHeadsHeadIDMaxSpeedParamsWithContext(ctx context.Context) *GetPrinterHeadsHeadIDMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDMaxSpeedParams{

		Context: ctx,
	}
}

// NewGetPrinterHeadsHeadIDMaxSpeedParamsWithHTTPClient creates a new GetPrinterHeadsHeadIDMaxSpeedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterHeadsHeadIDMaxSpeedParamsWithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDMaxSpeedParams{
		HTTPClient: client,
	}
}

/*GetPrinterHeadsHeadIDMaxSpeedParams contains all the parameters to send to the API endpoint
for the get printer heads head ID max speed operation typically these are written to a http.Request
*/
type GetPrinterHeadsHeadIDMaxSpeedParams struct {

	/*HeadID
	  ID of head of which to get the max speed of. Note that this speed also has a Z component. This api assumes that the head is the only part that moves.

	*/
	HeadID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) WithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDMaxSpeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) WithContext(ctx context.Context) *GetPrinterHeadsHeadIDMaxSpeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) WithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDMaxSpeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeadID adds the headID to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) WithHeadID(headID int64) *GetPrinterHeadsHeadIDMaxSpeedParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the get printer heads head ID max speed params
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterHeadsHeadIDMaxSpeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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