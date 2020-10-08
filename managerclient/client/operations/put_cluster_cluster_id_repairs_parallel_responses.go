// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-mgmt-commons/managerclient/models"
)

// PutClusterClusterIDRepairsParallelReader is a Reader for the PutClusterClusterIDRepairsParallel structure.
type PutClusterClusterIDRepairsParallelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterClusterIDRepairsParallelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutClusterClusterIDRepairsParallelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutClusterClusterIDRepairsParallelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutClusterClusterIDRepairsParallelOK creates a PutClusterClusterIDRepairsParallelOK with default headers values
func NewPutClusterClusterIDRepairsParallelOK() *PutClusterClusterIDRepairsParallelOK {
	return &PutClusterClusterIDRepairsParallelOK{}
}

/*PutClusterClusterIDRepairsParallelOK handles this case with default header values.

OK
*/
type PutClusterClusterIDRepairsParallelOK struct {
}

func (o *PutClusterClusterIDRepairsParallelOK) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repairs/parallel][%d] putClusterClusterIdRepairsParallelOK ", 200)
}

func (o *PutClusterClusterIDRepairsParallelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClusterClusterIDRepairsParallelDefault creates a PutClusterClusterIDRepairsParallelDefault with default headers values
func NewPutClusterClusterIDRepairsParallelDefault(code int) *PutClusterClusterIDRepairsParallelDefault {
	return &PutClusterClusterIDRepairsParallelDefault{
		_statusCode: code,
	}
}

/*PutClusterClusterIDRepairsParallelDefault handles this case with default header values.

Unexpected error
*/
type PutClusterClusterIDRepairsParallelDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the put cluster cluster ID repairs parallel default response
func (o *PutClusterClusterIDRepairsParallelDefault) Code() int {
	return o._statusCode
}

func (o *PutClusterClusterIDRepairsParallelDefault) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repairs/parallel][%d] PutClusterClusterIDRepairsParallel default  %+v", o._statusCode, o.Payload)
}

func (o *PutClusterClusterIDRepairsParallelDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDRepairsParallelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
