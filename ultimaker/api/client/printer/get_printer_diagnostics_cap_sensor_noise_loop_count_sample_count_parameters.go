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

// NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams creates a new GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams object
// with the default values initialized.
func NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams() *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParamsWithTimeout creates a new GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParamsWithTimeout(timeout time.Duration) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams{

		timeout: timeout,
	}
}

// NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParamsWithContext creates a new GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParamsWithContext(ctx context.Context) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams{

		Context: ctx,
	}
}

// NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParamsWithHTTPClient creates a new GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParamsWithHTTPClient(client *http.Client) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams{
		HTTPClient: client,
	}
}

/*GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams contains all the parameters to send to the API endpoint
for the get printer diagnostics cap sensor noise loop count sample count operation typically these are written to a http.Request
*/
type GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams struct {

	/*LoopCount
	  The number of times to get sample data

	*/
	LoopCount int32
	/*SampleCount
	  The number of times to get sample data

	*/
	SampleCount int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) WithTimeout(timeout time.Duration) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) WithContext(ctx context.Context) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) WithHTTPClient(client *http.Client) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLoopCount adds the loopCount to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) WithLoopCount(loopCount int32) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	o.SetLoopCount(loopCount)
	return o
}

// SetLoopCount adds the loopCount to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) SetLoopCount(loopCount int32) {
	o.LoopCount = loopCount
}

// WithSampleCount adds the sampleCount to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) WithSampleCount(sampleCount int32) *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams {
	o.SetSampleCount(sampleCount)
	return o
}

// SetSampleCount adds the sampleCount to the get printer diagnostics cap sensor noise loop count sample count params
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) SetSampleCount(sampleCount int32) {
	o.SampleCount = sampleCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterDiagnosticsCapSensorNoiseLoopCountSampleCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param loop_count
	if err := r.SetPathParam("loop_count", swag.FormatInt32(o.LoopCount)); err != nil {
		return err
	}

	// path param sample_count
	if err := r.SetPathParam("sample_count", swag.FormatInt32(o.SampleCount)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
