package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedReader is a Reader for the GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeed structure.
type GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK creates a GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK with default headers values
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK() *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK {
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK{}
}

/*GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK handles this case with default header values.

GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK get printer heads head Id extruders extruder Id feeder max speed o k
*/
type GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK struct {
	Payload float64
}

func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK) Error() string {
	return fmt.Sprintf("[GET /printer/heads/{head_id}/extruders/{extruder_id}/feeder/max_speed][%d] getPrinterHeadsHeadIdExtrudersExtruderIdFeederMaxSpeedOK  %+v", 200, o.Payload)
}

func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederMaxSpeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}