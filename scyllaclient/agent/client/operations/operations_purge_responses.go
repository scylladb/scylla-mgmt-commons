// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-mgmt-commons/scyllaclient/agent/models"
)

// OperationsPurgeReader is a Reader for the OperationsPurge structure.
type OperationsPurgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsPurgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsPurgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationsPurgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationsPurgeOK creates a OperationsPurgeOK with default headers values
func NewOperationsPurgeOK() *OperationsPurgeOK {
	return &OperationsPurgeOK{}
}

/*OperationsPurgeOK handles this case with default header values.

Job ID
*/
type OperationsPurgeOK struct {
	Payload *models.Jobid
	JobID   int64
}

func (o *OperationsPurgeOK) GetPayload() *models.Jobid {
	return o.Payload
}

func (o *OperationsPurgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewOperationsPurgeDefault creates a OperationsPurgeDefault with default headers values
func NewOperationsPurgeDefault(code int) *OperationsPurgeDefault {
	return &OperationsPurgeDefault{
		_statusCode: code,
	}
}

/*OperationsPurgeDefault handles this case with default header values.

Server error
*/
type OperationsPurgeDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the operations purge default response
func (o *OperationsPurgeDefault) Code() int {
	return o._statusCode
}

func (o *OperationsPurgeDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsPurgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *OperationsPurgeDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
