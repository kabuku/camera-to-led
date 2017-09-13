package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutSystemNameReader is a Reader for the PutSystemName structure.
type PutSystemNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSystemNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutSystemNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutSystemNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSystemNameNoContent creates a PutSystemNameNoContent with default headers values
func NewPutSystemNameNoContent() *PutSystemNameNoContent {
	return &PutSystemNameNoContent{}
}

/*PutSystemNameNoContent handles this case with default header values.

Name set
*/
type PutSystemNameNoContent struct {
}

func (o *PutSystemNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /system/name][%d] putSystemNameNoContent ", 204)
}

func (o *PutSystemNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemNameBadRequest creates a PutSystemNameBadRequest with default headers values
func NewPutSystemNameBadRequest() *PutSystemNameBadRequest {
	return &PutSystemNameBadRequest{}
}

/*PutSystemNameBadRequest handles this case with default header values.

Name is not set, because an invalid name is specified
*/
type PutSystemNameBadRequest struct {
}

func (o *PutSystemNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /system/name][%d] putSystemNameBadRequest ", 400)
}

func (o *PutSystemNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
