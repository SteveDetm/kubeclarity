// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/kubeclarity/api/client/models"
)

// GetRuntimeQuickscanConfigReader is a Reader for the GetRuntimeQuickscanConfig structure.
type GetRuntimeQuickscanConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeQuickscanConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeQuickscanConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRuntimeQuickscanConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeQuickscanConfigOK creates a GetRuntimeQuickscanConfigOK with default headers values
func NewGetRuntimeQuickscanConfigOK() *GetRuntimeQuickscanConfigOK {
	return &GetRuntimeQuickscanConfigOK{}
}

/* GetRuntimeQuickscanConfigOK describes a response with status code 200, with default header values.

Success
*/
type GetRuntimeQuickscanConfigOK struct {
	Payload *models.RuntimeQuickScanConfig
}

func (o *GetRuntimeQuickscanConfigOK) Error() string {
	return fmt.Sprintf("[GET /runtime/quickscan/config][%d] getRuntimeQuickscanConfigOK  %+v", 200, o.Payload)
}
func (o *GetRuntimeQuickscanConfigOK) GetPayload() *models.RuntimeQuickScanConfig {
	return o.Payload
}

func (o *GetRuntimeQuickscanConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeQuickScanConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeQuickscanConfigDefault creates a GetRuntimeQuickscanConfigDefault with default headers values
func NewGetRuntimeQuickscanConfigDefault(code int) *GetRuntimeQuickscanConfigDefault {
	return &GetRuntimeQuickscanConfigDefault{
		_statusCode: code,
	}
}

/* GetRuntimeQuickscanConfigDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetRuntimeQuickscanConfigDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get runtime quickscan config default response
func (o *GetRuntimeQuickscanConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeQuickscanConfigDefault) Error() string {
	return fmt.Sprintf("[GET /runtime/quickscan/config][%d] GetRuntimeQuickscanConfig default  %+v", o._statusCode, o.Payload)
}
func (o *GetRuntimeQuickscanConfigDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetRuntimeQuickscanConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
