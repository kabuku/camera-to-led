package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationReader is a Reader for the GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAcceleration structure.
type GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK creates a GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK with default headers values
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK() *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK {
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK{}
}

/*GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK handles this case with default header values.

GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK get printer heads head Id extruders extruder Id feeder acceleration o k
*/
type GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK struct {
	Payload float64
}

func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK) Error() string {
	return fmt.Sprintf("[GET /printer/heads/{head_id}/extruders/{extruder_id}/feeder/acceleration][%d] getPrinterHeadsHeadIdExtrudersExtruderIdFeederAccelerationOK  %+v", 200, o.Payload)
}

func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
