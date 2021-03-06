package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostPrinterLedBlinkReader is a Reader for the PostPrinterLedBlink structure.
type PostPrinterLedBlinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrinterLedBlinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostPrinterLedBlinkNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPrinterLedBlinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrinterLedBlinkNoContent creates a PostPrinterLedBlinkNoContent with default headers values
func NewPostPrinterLedBlinkNoContent() *PostPrinterLedBlinkNoContent {
	return &PostPrinterLedBlinkNoContent{}
}

/*PostPrinterLedBlinkNoContent handles this case with default header values.

blink set
*/
type PostPrinterLedBlinkNoContent struct {
}

func (o *PostPrinterLedBlinkNoContent) Error() string {
	return fmt.Sprintf("[POST /printer/led/blink][%d] postPrinterLedBlinkNoContent ", 204)
}

func (o *PostPrinterLedBlinkNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPrinterLedBlinkBadRequest creates a PostPrinterLedBlinkBadRequest with default headers values
func NewPostPrinterLedBlinkBadRequest() *PostPrinterLedBlinkBadRequest {
	return &PostPrinterLedBlinkBadRequest{}
}

/*PostPrinterLedBlinkBadRequest handles this case with default header values.

This is returned when frequency <= 0 or count <= 0 with a message
*/
type PostPrinterLedBlinkBadRequest struct {
}

func (o *PostPrinterLedBlinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /printer/led/blink][%d] postPrinterLedBlinkBadRequest ", 400)
}

func (o *PostPrinterLedBlinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
