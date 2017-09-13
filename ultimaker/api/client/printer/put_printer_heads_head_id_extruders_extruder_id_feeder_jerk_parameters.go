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

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams object
// with the default values initialized.
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams() *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParamsWithTimeout creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParamsWithTimeout(timeout time.Duration) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams{

		timeout: timeout,
	}
}

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParamsWithContext creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParamsWithContext(ctx context.Context) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams{

		Context: ctx,
	}
}

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParamsWithHTTPClient creates a new PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParamsWithHTTPClient(client *http.Client) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams{
		HTTPClient: client,
	}
}

/*PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams contains all the parameters to send to the API endpoint
for the put printer heads head ID extruders extruder ID feeder jerk operation typically these are written to a http.Request
*/
type PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams struct {

	/*ExtruderID
	  ID of extruder from which the feeder is fetched

	*/
	ExtruderID int64
	/*HeadID
	  ID of head from which the extruder is fetched

	*/
	HeadID int64
	/*Jerk
	  Target jerk

	*/
	Jerk *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) WithTimeout(timeout time.Duration) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) WithContext(ctx context.Context) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) WithHTTPClient(client *http.Client) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtruderID adds the extruderID to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) WithExtruderID(extruderID int64) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	o.SetExtruderID(extruderID)
	return o
}

// SetExtruderID adds the extruderId to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) SetExtruderID(extruderID int64) {
	o.ExtruderID = extruderID
}

// WithHeadID adds the headID to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) WithHeadID(headID int64) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WithJerk adds the jerk to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) WithJerk(jerk *float64) *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams {
	o.SetJerk(jerk)
	return o
}

// SetJerk adds the jerk to the put printer heads head ID extruders extruder ID feeder jerk params
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) SetJerk(jerk *float64) {
	o.Jerk = jerk
}

// WriteToRequest writes these params to a swagger request
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederJerkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Jerk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}