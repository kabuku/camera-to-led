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

// GetPrinterBedTemperatureReader is a Reader for the GetPrinterBedTemperature structure.
type GetPrinterBedTemperatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterBedTemperatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterBedTemperatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterBedTemperatureOK creates a GetPrinterBedTemperatureOK with default headers values
func NewGetPrinterBedTemperatureOK() *GetPrinterBedTemperatureOK {
	return &GetPrinterBedTemperatureOK{}
}

/*GetPrinterBedTemperatureOK handles this case with default header values.

Temperature of the bed
*/
type GetPrinterBedTemperatureOK struct {
	Payload *models.CurrentTargetNumberPair
}

func (o *GetPrinterBedTemperatureOK) Error() string {
	return fmt.Sprintf("[GET /printer/bed/temperature][%d] getPrinterBedTemperatureOK  %+v", 200, o.Payload)
}

func (o *GetPrinterBedTemperatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CurrentTargetNumberPair)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
