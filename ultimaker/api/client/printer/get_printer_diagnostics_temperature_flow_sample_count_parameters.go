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

// NewGetPrinterDiagnosticsTemperatureFlowSampleCountParams creates a new GetPrinterDiagnosticsTemperatureFlowSampleCountParams object
// with the default values initialized.
func NewGetPrinterDiagnosticsTemperatureFlowSampleCountParams() *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsTemperatureFlowSampleCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterDiagnosticsTemperatureFlowSampleCountParamsWithTimeout creates a new GetPrinterDiagnosticsTemperatureFlowSampleCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterDiagnosticsTemperatureFlowSampleCountParamsWithTimeout(timeout time.Duration) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsTemperatureFlowSampleCountParams{

		timeout: timeout,
	}
}

// NewGetPrinterDiagnosticsTemperatureFlowSampleCountParamsWithContext creates a new GetPrinterDiagnosticsTemperatureFlowSampleCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterDiagnosticsTemperatureFlowSampleCountParamsWithContext(ctx context.Context) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsTemperatureFlowSampleCountParams{

		Context: ctx,
	}
}

// NewGetPrinterDiagnosticsTemperatureFlowSampleCountParamsWithHTTPClient creates a new GetPrinterDiagnosticsTemperatureFlowSampleCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterDiagnosticsTemperatureFlowSampleCountParamsWithHTTPClient(client *http.Client) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	var ()
	return &GetPrinterDiagnosticsTemperatureFlowSampleCountParams{
		HTTPClient: client,
	}
}

/*GetPrinterDiagnosticsTemperatureFlowSampleCountParams contains all the parameters to send to the API endpoint
for the get printer diagnostics temperature flow sample count operation typically these are written to a http.Request
*/
type GetPrinterDiagnosticsTemperatureFlowSampleCountParams struct {

	/*Csv
	  If not zero, return the results as comma separated values instead of a normal json response.

	*/
	Csv *int32
	/*SampleCount
	  The amount of samples to get

	*/
	SampleCount int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) WithTimeout(timeout time.Duration) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) WithContext(ctx context.Context) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) WithHTTPClient(client *http.Client) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCsv adds the csv to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) WithCsv(csv *int32) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	o.SetCsv(csv)
	return o
}

// SetCsv adds the csv to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) SetCsv(csv *int32) {
	o.Csv = csv
}

// WithSampleCount adds the sampleCount to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) WithSampleCount(sampleCount int32) *GetPrinterDiagnosticsTemperatureFlowSampleCountParams {
	o.SetSampleCount(sampleCount)
	return o
}

// SetSampleCount adds the sampleCount to the get printer diagnostics temperature flow sample count params
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) SetSampleCount(sampleCount int32) {
	o.SampleCount = sampleCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Csv != nil {

		// query param csv
		var qrCsv int32
		if o.Csv != nil {
			qrCsv = *o.Csv
		}
		qCsv := swag.FormatInt32(qrCsv)
		if qCsv != "" {
			if err := r.SetQueryParam("csv", qCsv); err != nil {
				return err
			}
		}

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
