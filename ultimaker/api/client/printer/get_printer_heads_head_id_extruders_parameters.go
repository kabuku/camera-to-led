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

// NewGetPrinterHeadsHeadIDExtrudersParams creates a new GetPrinterHeadsHeadIDExtrudersParams object
// with the default values initialized.
func NewGetPrinterHeadsHeadIDExtrudersParams() *GetPrinterHeadsHeadIDExtrudersParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersParamsWithTimeout creates a new GetPrinterHeadsHeadIDExtrudersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterHeadsHeadIDExtrudersParamsWithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersParams{

		timeout: timeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersParamsWithContext creates a new GetPrinterHeadsHeadIDExtrudersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterHeadsHeadIDExtrudersParamsWithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersParams{

		Context: ctx,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersParamsWithHTTPClient creates a new GetPrinterHeadsHeadIDExtrudersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterHeadsHeadIDExtrudersParamsWithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersParams{
		HTTPClient: client,
	}
}

/*GetPrinterHeadsHeadIDExtrudersParams contains all the parameters to send to the API endpoint
for the get printer heads head ID extruders operation typically these are written to a http.Request
*/
type GetPrinterHeadsHeadIDExtrudersParams struct {

	/*HeadID
	  ID of head from which the extruders are fetched

	*/
	HeadID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) WithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) WithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) WithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeadID adds the headID to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) WithHeadID(headID int64) *GetPrinterHeadsHeadIDExtrudersParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the get printer heads head ID extruders params
func (o *GetPrinterHeadsHeadIDExtrudersParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterHeadsHeadIDExtrudersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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