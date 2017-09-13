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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kabuku/camera-to-led/ultimaker/api/models"
)

// NewPostPrinterLedBlinkParams creates a new PostPrinterLedBlinkParams object
// with the default values initialized.
func NewPostPrinterLedBlinkParams() *PostPrinterLedBlinkParams {
	var ()
	return &PostPrinterLedBlinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrinterLedBlinkParamsWithTimeout creates a new PostPrinterLedBlinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrinterLedBlinkParamsWithTimeout(timeout time.Duration) *PostPrinterLedBlinkParams {
	var ()
	return &PostPrinterLedBlinkParams{

		timeout: timeout,
	}
}

// NewPostPrinterLedBlinkParamsWithContext creates a new PostPrinterLedBlinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrinterLedBlinkParamsWithContext(ctx context.Context) *PostPrinterLedBlinkParams {
	var ()
	return &PostPrinterLedBlinkParams{

		Context: ctx,
	}
}

// NewPostPrinterLedBlinkParamsWithHTTPClient creates a new PostPrinterLedBlinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrinterLedBlinkParamsWithHTTPClient(client *http.Client) *PostPrinterLedBlinkParams {
	var ()
	return &PostPrinterLedBlinkParams{
		HTTPClient: client,
	}
}

/*PostPrinterLedBlinkParams contains all the parameters to send to the API endpoint
for the post printer led blink operation typically these are written to a http.Request
*/
type PostPrinterLedBlinkParams struct {

	/*Blink*/
	Blink *models.Blink

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post printer led blink params
func (o *PostPrinterLedBlinkParams) WithTimeout(timeout time.Duration) *PostPrinterLedBlinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post printer led blink params
func (o *PostPrinterLedBlinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post printer led blink params
func (o *PostPrinterLedBlinkParams) WithContext(ctx context.Context) *PostPrinterLedBlinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post printer led blink params
func (o *PostPrinterLedBlinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post printer led blink params
func (o *PostPrinterLedBlinkParams) WithHTTPClient(client *http.Client) *PostPrinterLedBlinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post printer led blink params
func (o *PostPrinterLedBlinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlink adds the blink to the post printer led blink params
func (o *PostPrinterLedBlinkParams) WithBlink(blink *models.Blink) *PostPrinterLedBlinkParams {
	o.SetBlink(blink)
	return o
}

// SetBlink adds the blink to the post printer led blink params
func (o *PostPrinterLedBlinkParams) SetBlink(blink *models.Blink) {
	o.Blink = blink
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrinterLedBlinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Blink == nil {
		o.Blink = new(models.Blink)
	}

	if err := r.SetBodyParam(o.Blink); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
