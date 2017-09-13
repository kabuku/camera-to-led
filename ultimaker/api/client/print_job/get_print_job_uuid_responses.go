package print_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kabuku/camera-to-led/ultimaker/api/models"
)

// GetPrintJobUUIDReader is a Reader for the GetPrintJobUUID structure.
type GetPrintJobUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrintJobUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrintJobUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrintJobUUIDOK creates a GetPrintJobUUIDOK with default headers values
func NewGetPrintJobUUIDOK() *GetPrintJobUUIDOK {
	return &GetPrintJobUUIDOK{}
}

/*GetPrintJobUUIDOK handles this case with default header values.

Unique identifier of this print job. In a UUID4 format.
*/
type GetPrintJobUUIDOK struct {
	Payload models.UUID
}

func (o *GetPrintJobUUIDOK) Error() string {
	return fmt.Sprintf("[GET /print_job/uuid][%d] getPrintJobUuidOK  %+v", 200, o.Payload)
}

func (o *GetPrintJobUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
