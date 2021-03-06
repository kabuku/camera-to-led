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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutBeepParams creates a new PutBeepParams object
// with the default values initialized.
func NewPutBeepParams() *PutBeepParams {
	var ()
	return &PutBeepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutBeepParamsWithTimeout creates a new PutBeepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutBeepParamsWithTimeout(timeout time.Duration) *PutBeepParams {
	var ()
	return &PutBeepParams{

		timeout: timeout,
	}
}

// NewPutBeepParamsWithContext creates a new PutBeepParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutBeepParamsWithContext(ctx context.Context) *PutBeepParams {
	var ()
	return &PutBeepParams{

		Context: ctx,
	}
}

// NewPutBeepParamsWithHTTPClient creates a new PutBeepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutBeepParamsWithHTTPClient(client *http.Client) *PutBeepParams {
	var ()
	return &PutBeepParams{
		HTTPClient: client,
	}
}

/*PutBeepParams contains all the parameters to send to the API endpoint
for the put beep operation typically these are written to a http.Request
*/
type PutBeepParams struct {

	/*Duration
	  The duration of the tone (in ms)

	*/
	Duration float64
	/*Frequency
	  The frequency of the tone (in Hz)

	*/
	Frequency float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put beep params
func (o *PutBeepParams) WithTimeout(timeout time.Duration) *PutBeepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put beep params
func (o *PutBeepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put beep params
func (o *PutBeepParams) WithContext(ctx context.Context) *PutBeepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put beep params
func (o *PutBeepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put beep params
func (o *PutBeepParams) WithHTTPClient(client *http.Client) *PutBeepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put beep params
func (o *PutBeepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDuration adds the duration to the put beep params
func (o *PutBeepParams) WithDuration(duration float64) *PutBeepParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the put beep params
func (o *PutBeepParams) SetDuration(duration float64) {
	o.Duration = duration
}

// WithFrequency adds the frequency to the put beep params
func (o *PutBeepParams) WithFrequency(frequency float64) *PutBeepParams {
	o.SetFrequency(frequency)
	return o
}

// SetFrequency adds the frequency to the put beep params
func (o *PutBeepParams) SetFrequency(frequency float64) {
	o.Frequency = frequency
}

// WriteToRequest writes these params to a swagger request
func (o *PutBeepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param duration
	qrDuration := o.Duration
	qDuration := swag.FormatFloat64(qrDuration)
	if qDuration != "" {
		if err := r.SetQueryParam("duration", qDuration); err != nil {
			return err
		}
	}

	// query param frequency
	qrFrequency := o.Frequency
	qFrequency := swag.FormatFloat64(qrFrequency)
	if qFrequency != "" {
		if err := r.SetQueryParam("frequency", qFrequency); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
