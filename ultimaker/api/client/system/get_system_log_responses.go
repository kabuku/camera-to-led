package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kabuku/camera-to-led/ultimaker/api/models"
)

// GetSystemLogReader is a Reader for the GetSystemLog structure.
type GetSystemLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSystemLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemLogOK creates a GetSystemLogOK with default headers values
func NewGetSystemLogOK() *GetSystemLogOK {
	return &GetSystemLogOK{}
}

/*GetSystemLogOK handles this case with default header values.

Log data
*/
type GetSystemLogOK struct {
	Payload models.SystemLog
}

func (o *GetSystemLogOK) Error() string {
	return fmt.Sprintf("[GET /system/log][%d] getSystemLogOK  %+v", 200, o.Payload)
}

func (o *GetSystemLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}