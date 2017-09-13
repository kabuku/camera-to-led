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

// GetPrinterHeadsHeadIDExtrudersExtruderIDHotendReader is a Reader for the GetPrinterHeadsHeadIDExtrudersExtruderIDHotend structure.
type GetPrinterHeadsHeadIDExtrudersExtruderIDHotendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK creates a GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK with default headers values
func NewGetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK() *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK {
	return &GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK{}
}

/*GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK handles this case with default header values.

GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK get printer heads head Id extruders extruder Id hotend o k
*/
type GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK struct {
	Payload *models.Hotend
}

func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK) Error() string {
	return fmt.Sprintf("[GET /printer/heads/{head_id}/extruders/{extruder_id}/hotend][%d] getPrinterHeadsHeadIdExtrudersExtruderIdHotendOK  %+v", 200, o.Payload)
}

func (o *GetPrinterHeadsHeadIDExtrudersExtruderIDHotendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hotend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}