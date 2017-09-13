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

// NewGetPrinterHeadsHeadIDPositionParams creates a new GetPrinterHeadsHeadIDPositionParams object
// with the default values initialized.
func NewGetPrinterHeadsHeadIDPositionParams() *GetPrinterHeadsHeadIDPositionParams {
	var ()
	return &GetPrinterHeadsHeadIDPositionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrinterHeadsHeadIDPositionParamsWithTimeout creates a new GetPrinterHeadsHeadIDPositionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrinterHeadsHeadIDPositionParamsWithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDPositionParams {
	var ()
	return &GetPrinterHeadsHeadIDPositionParams{

		timeout: timeout,
	}
}

// NewGetPrinterHeadsHeadIDPositionParamsWithContext creates a new GetPrinterHeadsHeadIDPositionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrinterHeadsHeadIDPositionParamsWithContext(ctx context.Context) *GetPrinterHeadsHeadIDPositionParams {
	var ()
	return &GetPrinterHeadsHeadIDPositionParams{

		Context: ctx,
	}
}

// NewGetPrinterHeadsHeadIDPositionParamsWithHTTPClient creates a new GetPrinterHeadsHeadIDPositionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrinterHeadsHeadIDPositionParamsWithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDPositionParams {
	var ()
	return &GetPrinterHeadsHeadIDPositionParams{
		HTTPClient: client,
	}
}

/*GetPrinterHeadsHeadIDPositionParams contains all the parameters to send to the API endpoint
for the get printer heads head ID position operation typically these are written to a http.Request
*/
type GetPrinterHeadsHeadIDPositionParams struct {

	/*HeadID
	  ID of head of which to get position. Note that this position also has a Z component. This api assumes that the head is the only part that moves.

	*/
	HeadID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) WithTimeout(timeout time.Duration) *GetPrinterHeadsHeadIDPositionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) WithContext(ctx context.Context) *GetPrinterHeadsHeadIDPositionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) WithHTTPClient(client *http.Client) *GetPrinterHeadsHeadIDPositionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeadID adds the headID to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) WithHeadID(headID int64) *GetPrinterHeadsHeadIDPositionParams {
	o.SetHeadID(headID)
	return o
}

// SetHeadID adds the headId to the get printer heads head ID position params
func (o *GetPrinterHeadsHeadIDPositionParams) SetHeadID(headID int64) {
	o.HeadID = headID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrinterHeadsHeadIDPositionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param head_id
	if err := r.SetPathParam("head_id", swag.FormatInt64(o.HeadID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}