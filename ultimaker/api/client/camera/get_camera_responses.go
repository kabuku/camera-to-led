package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kabuku/camera-to-led/ultimaker/api/models"
)

// GetCameraReader is a Reader for the GetCamera structure.
type GetCameraReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCameraReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCameraOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCameraOK creates a GetCameraOK with default headers values
func NewGetCameraOK() *GetCameraOK {
	return &GetCameraOK{}
}

/*GetCameraOK handles this case with default header values.

Camera object
*/
type GetCameraOK struct {
	Payload *models.Camera
}

func (o *GetCameraOK) Error() string {
	return fmt.Sprintf("[GET /camera][%d] getCameraOK  %+v", 200, o.Payload)
}

func (o *GetCameraOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Camera)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
