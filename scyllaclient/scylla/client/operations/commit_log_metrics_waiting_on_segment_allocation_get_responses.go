// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla/models"
)

// CommitLogMetricsWaitingOnSegmentAllocationGetReader is a Reader for the CommitLogMetricsWaitingOnSegmentAllocationGet structure.
type CommitLogMetricsWaitingOnSegmentAllocationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitLogMetricsWaitingOnSegmentAllocationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitLogMetricsWaitingOnSegmentAllocationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCommitLogMetricsWaitingOnSegmentAllocationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCommitLogMetricsWaitingOnSegmentAllocationGetOK creates a CommitLogMetricsWaitingOnSegmentAllocationGetOK with default headers values
func NewCommitLogMetricsWaitingOnSegmentAllocationGetOK() *CommitLogMetricsWaitingOnSegmentAllocationGetOK {
	return &CommitLogMetricsWaitingOnSegmentAllocationGetOK{}
}

/*CommitLogMetricsWaitingOnSegmentAllocationGetOK handles this case with default header values.

CommitLogMetricsWaitingOnSegmentAllocationGetOK commit log metrics waiting on segment allocation get o k
*/
type CommitLogMetricsWaitingOnSegmentAllocationGetOK struct {
}

func (o *CommitLogMetricsWaitingOnSegmentAllocationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitLogMetricsWaitingOnSegmentAllocationGetDefault creates a CommitLogMetricsWaitingOnSegmentAllocationGetDefault with default headers values
func NewCommitLogMetricsWaitingOnSegmentAllocationGetDefault(code int) *CommitLogMetricsWaitingOnSegmentAllocationGetDefault {
	return &CommitLogMetricsWaitingOnSegmentAllocationGetDefault{
		_statusCode: code,
	}
}

/*CommitLogMetricsWaitingOnSegmentAllocationGetDefault handles this case with default header values.

internal server error
*/
type CommitLogMetricsWaitingOnSegmentAllocationGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the commit log metrics waiting on segment allocation get default response
func (o *CommitLogMetricsWaitingOnSegmentAllocationGetDefault) Code() int {
	return o._statusCode
}

func (o *CommitLogMetricsWaitingOnSegmentAllocationGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CommitLogMetricsWaitingOnSegmentAllocationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CommitLogMetricsWaitingOnSegmentAllocationGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
