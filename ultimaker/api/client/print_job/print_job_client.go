package print_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new print job API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for print job API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPrintJob get print job API
*/
func (a *Client) GetPrintJob(params *GetPrintJobParams) (*GetPrintJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJob",
		Method:             "GET",
		PathPattern:        "/print_job",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobOK), nil

}

/*
GetPrintJobName get print job name API
*/
func (a *Client) GetPrintJobName(params *GetPrintJobNameParams) (*GetPrintJobNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobName",
		Method:             "GET",
		PathPattern:        "/print_job/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobNameOK), nil

}

/*
GetPrintJobProgress Get the (estimated) progress for the current print job, a value between 0 and 1
*/
func (a *Client) GetPrintJobProgress(params *GetPrintJobProgressParams) (*GetPrintJobProgressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobProgressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobProgress",
		Method:             "GET",
		PathPattern:        "/print_job/progress",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobProgressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobProgressOK), nil

}

/*
GetPrintJobSource From what source was the print job started. USB means it's started manually from the USB drive. WEB_API means it's being received by the WEB API. CALIBRATION_MENU means it's printing the XY offset print
*/
func (a *Client) GetPrintJobSource(params *GetPrintJobSourceParams) (*GetPrintJobSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobSource",
		Method:             "GET",
		PathPattern:        "/print_job/source",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobSourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobSourceOK), nil

}

/*
GetPrintJobSourceApplication If the origin equals to WEB_API, then this will return the application that sent the job
*/
func (a *Client) GetPrintJobSourceApplication(params *GetPrintJobSourceApplicationParams) (*GetPrintJobSourceApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobSourceApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobSourceApplication",
		Method:             "GET",
		PathPattern:        "/print_job/source_application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobSourceApplicationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobSourceApplicationOK), nil

}

/*
GetPrintJobSourceUser If the origin equals to WEB_API, then this will return the user who initiated the job
*/
func (a *Client) GetPrintJobSourceUser(params *GetPrintJobSourceUserParams) (*GetPrintJobSourceUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobSourceUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobSourceUser",
		Method:             "GET",
		PathPattern:        "/print_job/source_user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobSourceUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobSourceUserOK), nil

}

/*
GetPrintJobState Get the print job state
*/
func (a *Client) GetPrintJobState(params *GetPrintJobStateParams) (*GetPrintJobStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobState",
		Method:             "GET",
		PathPattern:        "/print_job/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobStateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobStateOK), nil

}

/*
GetPrintJobTimeElapsed Get the time elapsed (in seconds) since starting this print, including pauses etc.
*/
func (a *Client) GetPrintJobTimeElapsed(params *GetPrintJobTimeElapsedParams) (*GetPrintJobTimeElapsedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobTimeElapsedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobTimeElapsed",
		Method:             "GET",
		PathPattern:        "/print_job/time_elapsed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobTimeElapsedReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobTimeElapsedOK), nil

}

/*
GetPrintJobTimeTotal Get the (estimated) total time in seconds for this print, excluding pauses etc.
*/
func (a *Client) GetPrintJobTimeTotal(params *GetPrintJobTimeTotalParams) (*GetPrintJobTimeTotalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobTimeTotalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobTimeTotal",
		Method:             "GET",
		PathPattern:        "/print_job/time_total",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobTimeTotalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobTimeTotalOK), nil

}

/*
GetPrintJobUUID get print job UUID API
*/
func (a *Client) GetPrintJobUUID(params *GetPrintJobUUIDParams) (*GetPrintJobUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrintJobUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrintJobUUID",
		Method:             "GET",
		PathPattern:        "/print_job/uuid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrintJobUUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrintJobUUIDOK), nil

}

/*
PostPrintJob post print job API
*/
func (a *Client) PostPrintJob(params *PostPrintJobParams) (*PostPrintJobCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrintJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPrintJob",
		Method:             "POST",
		PathPattern:        "/print_job",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostPrintJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrintJobCreated), nil

}

/*
PutPrintJobState put print job state API
*/
func (a *Client) PutPrintJobState(params *PutPrintJobStateParams) (*PutPrintJobStateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPrintJobStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutPrintJobState",
		Method:             "PUT",
		PathPattern:        "/print_job/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutPrintJobStateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutPrintJobStateNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}