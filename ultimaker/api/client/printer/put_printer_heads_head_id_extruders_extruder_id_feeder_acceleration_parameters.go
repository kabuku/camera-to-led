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

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams object
// with the default values initialized.
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams() *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParamsWithTimeout creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParamsWithTimeout(timeout time.Duration) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams{

		timeout: timeout,
	}
}

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParamsWithContext creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParamsWithContext(ctx context.Context) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams{

		Context: ctx,
	}
}

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParamsWithHTTPClient creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParamsWithHTTPClient(client *http.Client) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams{
		HTTPClient: client,
	}
}

/*PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams contains all the parameters to send to the API endpoint
for the put printer heads head ID extruders extruder ID feeder acceleration operation typically these are written to a http.Request
*/
type PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams struct {

	/*Acceleration
	  Target acceleration speed

	*/
	Acceleration *float64
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

// WithTimeout adds the timeout to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) WithTimeout(timeout time.Duration) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) WithContext(ctx context.Context) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) WithHTTPClient(client *http.Client) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceleration adds the acceleration to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) WithAcceleration(acceleration *float64) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	o.SetAcceleration(acceleration)
	return o
}

// SetAcceleration adds the acceleration to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) SetAcceleration(acceleration *float64) {
	o.Acceleration = acceleration
}

// WithExtruderID adds the extruderID to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) WithExtruderID(extruderID int64) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	o.SetExtruderID(extruderID)
	return o
}

// SetExtruderID adds the extruderId to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) SetExtruderID(extruderID int64) {
	o.ExtruderID = extruderID
}

// WithHeadID adds the headID to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) WithHeadID(headID int64) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the put printer heads head ID extruders extruder ID feeder acceleration params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Acceleration); err != nil {
		return err
	}

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
