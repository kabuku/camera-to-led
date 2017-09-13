package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new camera API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for camera API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCamera Returns camera object
*/
func (a *Client) GetCamera(params *GetCameraParams) (*GetCameraOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCameraParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCamera",
		Method:             "GET",
		PathPattern:        "/camera",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCameraReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCameraOK), nil

}

/*
GetCameraFeed Get a link to the camera feed, this returns an url to a camera stream
*/
func (a *Client) GetCameraFeed(params *GetCameraFeedParams) (*GetCameraFeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCameraFeedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCameraFeed",
		Method:             "GET",
		PathPattern:        "/camera/feed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCameraFeedReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCameraFeedOK), nil

}

/*
GetCameraIndexSnapshot Get a redirection to the camera snapshot.
*/
func (a *Client) GetCameraIndexSnapshot(params *GetCameraIndexSnapshotParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCameraIndexSnapshotParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCameraIndexSnapshot",
		Method:             "GET",
		PathPattern:        "/camera/{index}/snapshot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCameraIndexSnapshotReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetCameraIndexStream Get a redirection to the camera live feed.
*/
func (a *Client) GetCameraIndexStream(params *GetCameraIndexStreamParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCameraIndexStreamParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCameraIndexStream",
		Method:             "GET",
		PathPattern:        "/camera/{index}/stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCameraIndexStreamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
