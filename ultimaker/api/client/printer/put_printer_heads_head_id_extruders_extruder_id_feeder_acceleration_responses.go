package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationReader is a Reader for the PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAcceleration structure.
type PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent creates a PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent with default headers values
func NewPutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent() *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent {
	return &PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent{}
}

/*PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent handles this case with default header values.

Acceleration set
*/
type PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent struct {
}

func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent) Error() string {
	return fmt.Sprintf("[PUT /printer/heads/{head_id}/extruders/{extruder_id}/feeder/acceleration][%d] putPrinterHeadsHeadIdExtrudersExtruderIdFeederAccelerationNoContent ", 204)
}

func (o *PutPrinterHeadsHeadIDExtrudersExtruderIDFeederAccelerationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
