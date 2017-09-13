package print_job

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
)

// NewPutPrintJobStateParams creates a new PutPrintJobStateParams object
// with the default values initialized.
func NewPutPrintJobStateParams() *PutPrintJobStateParams {
	var ()
	return &PutPrintJobStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPrintJobStateParamsWithTimeout creates a new PutPrintJobStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPrintJobStateParamsWithTimeout(timeout time.Duration) *PutPrintJobStateParams {
	var ()
	return &PutPrintJobStateParams{

		timeout: timeout,
	}
}

// NewPutPrintJobStateParamsWithContext creates a new PutPrintJobStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPrintJobStateParamsWithContext(ctx context.Context) *PutPrintJobStateParams {
	var ()
	return &PutPrintJobStateParams{

		Context: ctx,
	}
}

// NewPutPrintJobStateParamsWithHTTPClient creates a new PutPrintJobStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPrintJobStateParamsWithHTTPClient(client *http.Client) *PutPrintJobStateParams {
	var ()
	return &PutPrintJobStateParams{
		HTTPClient: client,
	}
}

/*PutPrintJobStateParams contains all the parameters to send to the API endpoint
for the put print job state operation typically these are written to a http.Request
*/
type PutPrintJobStateParams struct {

	/*Target
	  "print", "pause" or "abort". Change the current state of the print. Note that only changes to abort / pause are always allowed and changing to print only when state is paused.

	*/
	Target *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put print job state params
func (o *PutPrintJobStateParams) WithTimeout(timeout time.Duration) *PutPrintJobStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put print job state params
func (o *PutPrintJobStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put print job state params
func (o *PutPrintJobStateParams) WithContext(ctx context.Context) *PutPrintJobStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put print job state params
func (o *PutPrintJobStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put print job state params
func (o *PutPrintJobStateParams) WithHTTPClient(client *http.Client) *PutPrintJobStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put print job state params
func (o *PutPrintJobStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTarget adds the target to the put print job state params
func (o *PutPrintJobStateParams) WithTarget(target *string) *PutPrintJobStateParams {
	o.SetTarget(target)
	return o
}

// SetTarget adds the target to the put print job state params
func (o *PutPrintJobStateParams) SetTarget(target *string) {
	o.Target = target
}

// WriteToRequest writes these params to a swagger request
func (o *PutPrintJobStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Target); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}