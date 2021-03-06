package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutPrinterBedPreHeatReader is a Reader for the PutPrinterBedPreHeat structure.
type PutPrinterBedPreHeatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPrinterBedPreHeatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutPrinterBedPreHeatNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutPrinterBedPreHeatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPrinterBedPreHeatNoContent creates a PutPrinterBedPreHeatNoContent with default headers values
func NewPutPrinterBedPreHeatNoContent() *PutPrinterBedPreHeatNoContent {
	return &PutPrinterBedPreHeatNoContent{}
}

/*PutPrinterBedPreHeatNoContent handles this case with default header values.

Preheating Temperature set
*/
type PutPrinterBedPreHeatNoContent struct {
}

func (o *PutPrinterBedPreHeatNoContent) Error() string {
	return fmt.Sprintf("[PUT /printer/bed/pre_heat][%d] putPrinterBedPreHeatNoContent ", 204)
}

func (o *PutPrinterBedPreHeatNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutPrinterBedPreHeatBadRequest creates a PutPrinterBedPreHeatBadRequest with default headers values
func NewPutPrinterBedPreHeatBadRequest() *PutPrinterBedPreHeatBadRequest {
	return &PutPrinterBedPreHeatBadRequest{}
}

/*PutPrinterBedPreHeatBadRequest handles this case with default header values.

Bad request (invalid parameters)
*/
type PutPrinterBedPreHeatBadRequest struct {
}

func (o *PutPrinterBedPreHeatBadRequest) Error() string {
	return fmt.Sprintf("[PUT /printer/bed/pre_heat][%d] putPrinterBedPreHeatBadRequest ", 400)
}

func (o *PutPrinterBedPreHeatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
