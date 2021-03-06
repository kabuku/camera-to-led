package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kabuku/camera-to-led/ultimaker/api/models"
)

// GetPrinterHeadsHeadIDReader is a Reader for the GetPrinterHeadsHeadID structure.
type GetPrinterHeadsHeadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterHeadsHeadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterHeadsHeadIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetPrinterHeadsHeadIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterHeadsHeadIDOK creates a GetPrinterHeadsHeadIDOK with default headers values
func NewGetPrinterHeadsHeadIDOK() *GetPrinterHeadsHeadIDOK {
	return &GetPrinterHeadsHeadIDOK{}
}

/*GetPrinterHeadsHeadIDOK handles this case with default header values.

GetPrinterHeadsHeadIDOK get printer heads head Id o k
*/
type GetPrinterHeadsHeadIDOK struct {
	Payload *models.Head
}

func (o *GetPrinterHeadsHeadIDOK) Error() string {
	return fmt.Sprintf("[GET /printer/heads/{head_id}][%d] getPrinterHeadsHeadIdOK  %+v", 200, o.Payload)
}

func (o *GetPrinterHeadsHeadIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Head)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrinterHeadsHeadIDNotFound creates a GetPrinterHeadsHeadIDNotFound with default headers values
func NewGetPrinterHeadsHeadIDNotFound() *GetPrinterHeadsHeadIDNotFound {
	return &GetPrinterHeadsHeadIDNotFound{}
}

/*GetPrinterHeadsHeadIDNotFound handles this case with default header values.

Head was not found. Note that this means that all deeper (eg: getting position, extruders, etc.) calls will also return a 404
*/
type GetPrinterHeadsHeadIDNotFound struct {
}

func (o *GetPrinterHeadsHeadIDNotFound) Error() string {
	return fmt.Sprintf("[GET /printer/heads/{head_id}][%d] getPrinterHeadsHeadIdNotFound ", 404)
}

func (o *GetPrinterHeadsHeadIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
