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

// GetPrinterReader is a Reader for the GetPrinter structure.
type GetPrinterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterOK creates a GetPrinterOK with default headers values
func NewGetPrinterOK() *GetPrinterOK {
	return &GetPrinterOK{}
}

/*GetPrinterOK handles this case with default header values.

Printer object
*/
type GetPrinterOK struct {
	Payload *models.Printer
}

func (o *GetPrinterOK) Error() string {
	return fmt.Sprintf("[GET /printer][%d] getPrinterOK  %+v", 200, o.Payload)
}

func (o *GetPrinterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Printer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
