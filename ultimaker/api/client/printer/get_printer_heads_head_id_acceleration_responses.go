package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPrinterHeadsHeadIDAccelerationReader is a Reader for the GetPrinterHeadsHeadIDAcceleration structure.
type GetPrinterHeadsHeadIDAccelerationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterHeadsHeadIDAccelerationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterHeadsHeadIDAccelerationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterHeadsHeadIDAccelerationOK creates a GetPrinterHeadsHeadIDAccelerationOK with default headers values
func NewGetPrinterHeadsHeadIDAccelerationOK() *GetPrinterHeadsHeadIDAccelerationOK {
	return &GetPrinterHeadsHeadIDAccelerationOK{}
}

/*GetPrinterHeadsHeadIDAccelerationOK handles this case with default header values.

GetPrinterHeadsHeadIDAccelerationOK get printer heads head Id acceleration o k
*/
type GetPrinterHeadsHeadIDAccelerationOK struct {
	Payload float64
}

func (o *GetPrinterHeadsHeadIDAccelerationOK) Error() string {
	return fmt.Sprintf("[GET /printer/heads/{head_id}/acceleration][%d] getPrinterHeadsHeadIdAccelerationOK  %+v", 200, o.Payload)
}

func (o *GetPrinterHeadsHeadIDAccelerationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
