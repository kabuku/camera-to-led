package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAuthCheckID Check if the given ID is authorized for printer access. Will return 'authorized' when the end user has selected that this application is allowed to use the printer. Will return 'unauthorized' when the user has selected that the application is not allowed to access the printer. Will return 'unknown' when the end user has not selected any option yet.
*/
func (a *Client) GetAuthCheckID(params *GetAuthCheckIDParams) (*GetAuthCheckIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthCheckIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAuthCheckID",
		Method:             "GET",
		PathPattern:        "/auth/check/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAuthCheckIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuthCheckIDOK), nil

}

/*
GetAuthVerify This API call always does authentication checking for digest authentication. Invalid digest id/key combinations will generate a 401 result.
*/
func (a *Client) GetAuthVerify(params *GetAuthVerifyParams) (*GetAuthVerifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthVerifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAuthVerify",
		Method:             "GET",
		PathPattern:        "/auth/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAuthVerifyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuthVerifyOK), nil

}

/*
PostAuthRequest Request authentication from the printer. This generates new id/key combination that has to be used as username/password in the digest authentication on certain APIs.
*/
func (a *Client) PostAuthRequest(params *PostAuthRequestParams) (*PostAuthRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAuthRequest",
		Method:             "POST",
		PathPattern:        "/auth/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAuthRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAuthRequestOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
