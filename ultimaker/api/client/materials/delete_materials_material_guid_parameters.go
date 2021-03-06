package materials

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

// NewDeleteMaterialsMaterialGUIDParams creates a new DeleteMaterialsMaterialGUIDParams object
// with the default values initialized.
func NewDeleteMaterialsMaterialGUIDParams() *DeleteMaterialsMaterialGUIDParams {
	var ()
	return &DeleteMaterialsMaterialGUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMaterialsMaterialGUIDParamsWithTimeout creates a new DeleteMaterialsMaterialGUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMaterialsMaterialGUIDParamsWithTimeout(timeout time.Duration) *DeleteMaterialsMaterialGUIDParams {
	var ()
	return &DeleteMaterialsMaterialGUIDParams{

		timeout: timeout,
	}
}

// NewDeleteMaterialsMaterialGUIDParamsWithContext creates a new DeleteMaterialsMaterialGUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMaterialsMaterialGUIDParamsWithContext(ctx context.Context) *DeleteMaterialsMaterialGUIDParams {
	var ()
	return &DeleteMaterialsMaterialGUIDParams{

		Context: ctx,
	}
}

// NewDeleteMaterialsMaterialGUIDParamsWithHTTPClient creates a new DeleteMaterialsMaterialGUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMaterialsMaterialGUIDParamsWithHTTPClient(client *http.Client) *DeleteMaterialsMaterialGUIDParams {
	var ()
	return &DeleteMaterialsMaterialGUIDParams{
		HTTPClient: client,
	}
}

/*DeleteMaterialsMaterialGUIDParams contains all the parameters to send to the API endpoint
for the delete materials material GUID operation typically these are written to a http.Request
*/
type DeleteMaterialsMaterialGUIDParams struct {

	/*MaterialGUID
	  GUID of material to delete

	*/
	MaterialGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) WithTimeout(timeout time.Duration) *DeleteMaterialsMaterialGUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) WithContext(ctx context.Context) *DeleteMaterialsMaterialGUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) WithHTTPClient(client *http.Client) *DeleteMaterialsMaterialGUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaterialGUID adds the materialGUID to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) WithMaterialGUID(materialGUID string) *DeleteMaterialsMaterialGUIDParams {
	o.SetMaterialGUID(materialGUID)
	return o
}

// SetMaterialGUID adds the materialGuid to the delete materials material GUID params
func (o *DeleteMaterialsMaterialGUIDParams) SetMaterialGUID(materialGUID string) {
	o.MaterialGUID = materialGUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMaterialsMaterialGUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param material_guid
	if err := r.SetPathParam("material_guid", o.MaterialGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
