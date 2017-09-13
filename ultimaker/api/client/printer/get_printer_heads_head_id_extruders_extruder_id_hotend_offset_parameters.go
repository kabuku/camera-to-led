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

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams object
// with the default values initialized.
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams() *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParamsWithTimeout creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParamsWithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams{

		timeout: timeout,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParamsWithContext creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParamsWithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams{

		Context: ctx,
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParamsWithHTTPClient creates a new GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParamsWithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	var ()
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams{
		HTTPClient: client,
	}
}

/*GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams contains all the parameters to send to the API endpoint
for the get printer heads head ID extruders extruder ID hotend offset operation typically these are written to a http.Request
*/
type GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams struct {

	/*ExtruderID
	  ID of extruder to fetch

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

// WithTimeout adds the timeout to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) WithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) WithContext(ctx context.Context) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) WithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtruderID adds the extruderID to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) WithExtruderID(extruderID int64) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	o.SetExtruderID(extruderID)
	return o
}

// SetExtruderID adds the extruderId to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) SetExtruderID(extruderID int64) {
	o.ExtruderID = extruderID
}

// WithHeadID adds the headID to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) WithHeadID(headID int64) *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the get printer heads head ID extruders extruder ID hotend offset params
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOffsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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