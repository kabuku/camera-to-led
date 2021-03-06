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

	"github.com/kabuku/camera-to-led/ultimaker/api/models"
)

// NewPutPrinterHeadsHeadIDJerkParams creates a new PutPrinterHeadsHeadIDJerkParams object
// with the default values initialized.
func NewPutPrinterHeadsHeadIDJerkParams() *PutPrinterHeadsHeadIDJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDJerkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPrinterHeadsHeadIDJerkParamsWithTimeout creates a new PutPrinterHeadsHeadIDJerkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPrinterHeadsHeadIDJerkParamsWithTimeout(timeout time.Duration) *PutPrinterHeadsHeadIDJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDJerkParams{

		timeout: timeout,
	}
}

// NewPutPrinterHeadsHeadIDJerkParamsWithContext creates a new PutPrinterHeadsHeadIDJerkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPrinterHeadsHeadIDJerkParamsWithContext(ctx context.Context) *PutPrinterHeadsHeadIDJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDJerkParams{

		Context: ctx,
	}
}

// NewPutPrinterHeadsHeadIDJerkParamsWithHTTPClient creates a new PutPrinterHeadsHeadIDJerkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPrinterHeadsHeadIDJerkParamsWithHTTPClient(client *http.Client) *PutPrinterHeadsHeadIDJerkParams {
	var ()
	return &PutPrinterHeadsHeadIDJerkParams{
		HTTPClient: client,
	}
}

/*PutPrinterHeadsHeadIDJerkParams contains all the parameters to send to the API endpoint
for the put printer heads head ID jerk operation typically these are written to a http.Request
*/
type PutPrinterHeadsHeadIDJerkParams struct {

	/*HeadID*/
	HeadID int64
	/*Jerk
	  Target jerk

	*/
	Jerk *models.XYZ

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) WithTimeout(timeout time.Duration) *PutPrinterHeadsHeadIDJerkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) WithContext(ctx context.Context) *PutPrinterHeadsHeadIDJerkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) WithHTTPClient(client *http.Client) *PutPrinterHeadsHeadIDJerkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeadID adds the headID to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) WithHeadID(headID int64) *PutPrinterHeadsHeadIDJerkParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WithJerk adds the jerk to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) WithJerk(jerk *models.XYZ) *PutPrinterHeadsHeadIDJerkParams {
	o.SetJerk(jerk)
	return o
}

// SetJerk adds the jerk to the put printer heads head ID jerk params
func (o *PutPrinterHeadsHeadIDJerkParams) SetJerk(jerk *models.XYZ) {
	o.Jerk = jerk
}

// WriteToRequest writes these params to a swagger request
func (o *PutPrinterHeadsHeadIDJerkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param head_id
	if err := r.SetPathParam("head_id", swag.FormatInt64(o.HeadID)); err != nil {
		return err
	}

	if o.Jerk == nil {
		o.Jerk = new(models.XYZ)
	}

	if err := r.SetBodyParam(o.Jerk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
