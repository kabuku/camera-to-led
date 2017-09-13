package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutPrinterHeadsHeadIDMaxSpeedReader is a Reader for the PutPrinterHeadsHeadIDMaxSpeed structure.
type PutPrinterHeadsHeadIDMaxSpeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPrinterHeadsHeadIDMaxSpeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutPrinterHeadsHeadIDMaxSpeedNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPrinterHeadsHeadIDMaxSpeedNoContent creates a PutPrinterHeadsHeadIDMaxSpeedNoContent with default headers values
func NewPutPrinterHeadsHeadIDMaxSpeedNoContent() *PutPrinterHeadsHeadIDMaxSpeedNoContent {
	return &PutPrinterHeadsHeadIDMaxSpeedNoContent{}
}

/*PutPrinterHeadsHeadIDMaxSpeedNoContent handles this case with default header values.

Max speed set
*/
type PutPrinterHeadsHeadIDMaxSpeedNoContent struct {
}

func (o *PutPrinterHeadsHeadIDMaxSpeedNoContent) Error() string {
	return fmt.Sprintf("[PUT /printer/heads/{head_id}/max_speed][%d] putPrinterHeadsHeadIdMaxSpeedNoContent ", 204)
}

func (o *PutPrinterHeadsHeadIDMaxSpeedNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
