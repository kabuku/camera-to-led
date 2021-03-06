package camera

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

// NewGetCameraIndexSnapshotParams creates a new GetCameraIndexSnapshotParams object
// with the default values initialized.
func NewGetCameraIndexSnapshotParams() *GetCameraIndexSnapshotParams {
	var ()
	return &GetCameraIndexSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCameraIndexSnapshotParamsWithTimeout creates a new GetCameraIndexSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCameraIndexSnapshotParamsWithTimeout(timeout time.Duration) *GetCameraIndexSnapshotParams {
	var ()
	return &GetCameraIndexSnapshotParams{

		timeout: timeout,
	}
}

// NewGetCameraIndexSnapshotParamsWithContext creates a new GetCameraIndexSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCameraIndexSnapshotParamsWithContext(ctx context.Context) *GetCameraIndexSnapshotParams {
	var ()
	return &GetCameraIndexSnapshotParams{

		Context: ctx,
	}
}

// NewGetCameraIndexSnapshotParamsWithHTTPClient creates a new GetCameraIndexSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCameraIndexSnapshotParamsWithHTTPClient(client *http.Client) *GetCameraIndexSnapshotParams {
	var ()
	return &GetCameraIndexSnapshotParams{
		HTTPClient: client,
	}
}

/*GetCameraIndexSnapshotParams contains all the parameters to send to the API endpoint
for the get camera index snapshot operation typically these are written to a http.Request
*/
type GetCameraIndexSnapshotParams struct {

	/*Index
	  index of the camera to get the snapshot from.

	*/
	Index float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) WithTimeout(timeout time.Duration) *GetCameraIndexSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) WithContext(ctx context.Context) *GetCameraIndexSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) WithHTTPClient(client *http.Client) *GetCameraIndexSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) WithIndex(index float64) *GetCameraIndexSnapshotParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get camera index snapshot params
func (o *GetCameraIndexSnapshotParams) SetIndex(index float64) {
	o.Index = index
}

// WriteToRequest writes these params to a swagger request
func (o *GetCameraIndexSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatFloat64(o.Index)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
