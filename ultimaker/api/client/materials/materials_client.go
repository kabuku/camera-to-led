package materials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new materials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for materials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteMaterialsMaterialGUID delete materials material GUID API
*/
func (a *Client) DeleteMaterialsMaterialGUID(params *DeleteMaterialsMaterialGUIDParams) (*DeleteMaterialsMaterialGUIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMaterialsMaterialGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteMaterialsMaterialGUID",
		Method:             "DELETE",
		PathPattern:        "/materials/{material_guid}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteMaterialsMaterialGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMaterialsMaterialGUIDNoContent), nil

}

/*
GetMaterials get materials API
*/
func (a *Client) GetMaterials(params *GetMaterialsParams) (*GetMaterialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMaterialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMaterials",
		Method:             "GET",
		PathPattern:        "/materials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMaterialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMaterialsOK), nil

}

/*
GetMaterialsMaterialGUID get materials material GUID API
*/
func (a *Client) GetMaterialsMaterialGUID(params *GetMaterialsMaterialGUIDParams) (*GetMaterialsMaterialGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMaterialsMaterialGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMaterialsMaterialGUID",
		Method:             "GET",
		PathPattern:        "/materials/{material_guid}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMaterialsMaterialGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMaterialsMaterialGUIDOK), nil

}

/*
PostMaterials post materials API
*/
func (a *Client) PostMaterials(params *PostMaterialsParams) (*PostMaterialsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMaterialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostMaterials",
		Method:             "POST",
		PathPattern:        "/materials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostMaterialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostMaterialsNoContent), nil

}

/*
PutMaterialsMaterialGUID put materials material GUID API
*/
func (a *Client) PutMaterialsMaterialGUID(params *PutMaterialsMaterialGUIDParams) (*PutMaterialsMaterialGUIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMaterialsMaterialGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutMaterialsMaterialGUID",
		Method:             "PUT",
		PathPattern:        "/materials/{material_guid}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutMaterialsMaterialGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutMaterialsMaterialGUIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
