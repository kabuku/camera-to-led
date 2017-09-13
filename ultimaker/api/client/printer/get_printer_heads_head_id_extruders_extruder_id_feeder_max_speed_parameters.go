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

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams object
// with the default values initialized.
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams() *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParamsWithTimeout creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParamsWithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams{

		timeout: timeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParamsWithContext creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParamsWithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams{

		Context: ctx,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParamsWithHTTPClient creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParamsWithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams{
		HTTPClient: client,
	}
}

/*GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams contains all the parameters to send to the API endpoint
for the get printer heads head ID extruders extruder ID feeder max speed operation typically these are written to a http.Request
*/
type GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams struct {

	/*ExtruderID
	  ID of extruder from which the feeder is fetched

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

// WithTimeout adds the timeout to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) WithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) WithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) WithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtruderID adds the extruderID to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) WithExtruderID(extruderID int64) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	o.SetExtruderID(extruderID)
	return o
}

// SetExtruderID adds the extruderId to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) SetExtruderID(extruderID int64) {
	o.ExtruderID = extruderID
}

// WithHeadID adds the headID to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) WithHeadID(headID int64) *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the get printer heads head ID extruders extruder ID feeder max speed params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
