package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetSystemNameReader is a Reader for the GetSystemName structure.
type GetSystemNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSystemNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemNameOK creates a GetSystemNameOK with default headers values
func NewGetSystemNameOK() *GetSystemNameOK {
	return &GetSystemNameOK{}
}

/*GetSystemNameOK handles this case with default header values.

name
*/
type GetSystemNameOK struct {
	Payload string
}

func (o *GetSystemNameOK) Error() string {
	return fmt.Sprintf("[GET /system/name][%d] getSystemNameOK  %+v", 200, o.Payload)
}

func (o *GetSystemNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}