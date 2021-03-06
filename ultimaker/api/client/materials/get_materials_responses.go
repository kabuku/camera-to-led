package materials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetMaterialsReader is a Reader for the GetMaterials structure.
type GetMaterialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMaterialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMaterialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMaterialsOK creates a GetMaterialsOK with default headers values
func NewGetMaterialsOK() *GetMaterialsOK {
	return &GetMaterialsOK{}
}

/*GetMaterialsOK handles this case with default header values.

All known material XML files, one string for each material.
*/
type GetMaterialsOK struct {
	Payload []string
}

func (o *GetMaterialsOK) Error() string {
	return fmt.Sprintf("[GET /materials][%d] getMaterialsOK  %+v", 200, o.Payload)
}

func (o *GetMaterialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
