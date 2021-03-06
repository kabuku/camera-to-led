package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCameraIndexSnapshotReader is a Reader for the GetCameraIndexSnapshot structure.
type GetCameraIndexSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCameraIndexSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 302:
		result := NewGetCameraIndexSnapshotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCameraIndexSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCameraIndexSnapshotFound creates a GetCameraIndexSnapshotFound with default headers values
func NewGetCameraIndexSnapshotFound() *GetCameraIndexSnapshotFound {
	return &GetCameraIndexSnapshotFound{}
}

/*GetCameraIndexSnapshotFound handles this case with default header values.

Redirection to the camera snapshot.
*/
type GetCameraIndexSnapshotFound struct {
}

func (o *GetCameraIndexSnapshotFound) Error() string {
	return fmt.Sprintf("[GET /camera/{index}/snapshot][%d] getCameraIndexSnapshotFound ", 302)
}

func (o *GetCameraIndexSnapshotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCameraIndexSnapshotNotFound creates a GetCameraIndexSnapshotNotFound with default headers values
func NewGetCameraIndexSnapshotNotFound() *GetCameraIndexSnapshotNotFound {
	return &GetCameraIndexSnapshotNotFound{}
}

/*GetCameraIndexSnapshotNotFound handles this case with default header values.

Camera with this index is not available in the system.
*/
type GetCameraIndexSnapshotNotFound struct {
}

func (o *GetCameraIndexSnapshotNotFound) Error() string {
	return fmt.Sprintf("[GET /camera/{index}/snapshot][%d] getCameraIndexSnapshotNotFound ", 404)
}

func (o *GetCameraIndexSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
