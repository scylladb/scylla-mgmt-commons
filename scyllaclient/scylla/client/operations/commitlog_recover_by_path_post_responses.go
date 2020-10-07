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

// CommitlogRecoverByPathPostReader is a Reader for the CommitlogRecoverByPathPost structure.
type CommitlogRecoverByPathPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitlogRecoverByPathPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitlogRecoverByPathPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCommitlogRecoverByPathPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCommitlogRecoverByPathPostOK creates a CommitlogRecoverByPathPostOK with default headers values
func NewCommitlogRecoverByPathPostOK() *CommitlogRecoverByPathPostOK {
	return &CommitlogRecoverByPathPostOK{}
}

/*CommitlogRecoverByPathPostOK handles this case with default header values.

CommitlogRecoverByPathPostOK commitlog recover by path post o k
*/
type CommitlogRecoverByPathPostOK struct {
}

func (o *CommitlogRecoverByPathPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitlogRecoverByPathPostDefault creates a CommitlogRecoverByPathPostDefault with default headers values
func NewCommitlogRecoverByPathPostDefault(code int) *CommitlogRecoverByPathPostDefault {
	return &CommitlogRecoverByPathPostDefault{
		_statusCode: code,
	}
}

/*CommitlogRecoverByPathPostDefault handles this case with default header values.

internal server error
*/
type CommitlogRecoverByPathPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the commitlog recover by path post default response
func (o *CommitlogRecoverByPathPostDefault) Code() int {
	return o._statusCode
}

func (o *CommitlogRecoverByPathPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CommitlogRecoverByPathPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CommitlogRecoverByPathPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
