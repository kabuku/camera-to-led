package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPrinterDiagnosticsTemperatureFlowSampleCountReader is a Reader for the GetPrinterDiagnosticsTemperatureFlowSampleCount structure.
type GetPrinterDiagnosticsTemperatureFlowSampleCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrinterDiagnosticsTemperatureFlowSampleCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterDiagnosticsTemperatureFlowSampleCountOK creates a GetPrinterDiagnosticsTemperatureFlowSampleCountOK with default headers values
func NewGetPrinterDiagnosticsTemperatureFlowSampleCountOK() *GetPrinterDiagnosticsTemperatureFlowSampleCountOK {
	return &GetPrinterDiagnosticsTemperatureFlowSampleCountOK{}
}

/*GetPrinterDiagnosticsTemperatureFlowSampleCountOK handles this case with default header values.

A 2 dimensional array of sample data. First row of the array contains names of each column. All the other rows contain the actual sample data.
*/
type GetPrinterDiagnosticsTemperatureFlowSampleCountOK struct {
}

func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountOK) Error() string {
	return fmt.Sprintf("[GET /printer/diagnostics/temperature_flow/{sample_count}][%d] getPrinterDiagnosticsTemperatureFlowSampleCountOK ", 200)
}

func (o *GetPrinterDiagnosticsTemperatureFlowSampleCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
