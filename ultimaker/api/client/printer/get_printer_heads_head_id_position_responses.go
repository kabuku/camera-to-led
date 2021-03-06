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

// GetPrinterHeadsHeadIDPositionReader is a Reader for the GetPrinterHeadsHeadIDPosition structure.
type GetPrinterHeadsHeadIDPositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterHeadsHeadIDPositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterHeadsHeadIDPositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterHeadsHeadIDPositionOK creates a GetPrinterHeadsHeadIDPositionOK with default headers values
func NewGetPrinterHeadsHeadIDPositionOK() *GetPrinterHeadsHeadIDPositionOK {
	return &GetPrinterHeadsHeadIDPositionOK{}
}

/*GetPrinterHeadsHeadIDPositionOK handles this case with default header values.

GetPrinterHeadsHeadIDPositionOK get printer heads head Id position o k
*/
type GetPrinterHeadsHeadIDPositionOK struct {
	Payload *models.XYZ
}

func (o *GetPrinterHeadsHeadIDPositionOK) Error() string {
	return fmt.Sprintf("[GET /printer/heads/{head_id}/position][%d] getPrinterHeadsHeadIdPositionOK  %+v", 200, o.Payload)
}

func (o *GetPrinterHeadsHeadIDPositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.XYZ)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
