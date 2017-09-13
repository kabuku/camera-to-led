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

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams object
// with the default values initialized.
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams() *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParamsWithTimeout creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParamsWithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams{

		timeout: timeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParamsWithContext creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParamsWithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams{

		Context: ctx,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParamsWithHTTPClient creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParamsWithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams{
		HTTPClient: client,
	}
}

/*GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams contains all the parameters to send to the API endpoint
for the get printer heads head ID extruders extruder ID active material operation typically these are written to a http.Request
*/
type GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams struct {

	/*ExtruderID
	  ID of extruder

	*/
	ExtruderID int64
	/*HeadID
	  ID of head from which the extruder is fetched

	*/
	HeadID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) WithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) WithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) WithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtruderID adds the extruderID to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) WithExtruderID(extruderID int64) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	o.SetExtruderID(extruderID)
	return o
}

// SetExtruderID adds the extruderId to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) SetExtruderID(extruderID int64) {
	o.ExtruderID = extruderID
}

// WithHeadID adds the headID to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) WithHeadID(headID int64) *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the get printer heads head ID extruders extruder ID active material params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDActiveMaterialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extruder_id
	if err := r.SetPathParam("extruder_id", swag.FormatInt64(o.ExtruderID)); err != nil {
		return err
	}

	// path param head_id
	if err := r.SetPathParam("head_id", swag.FormatInt64(o.HeadID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
