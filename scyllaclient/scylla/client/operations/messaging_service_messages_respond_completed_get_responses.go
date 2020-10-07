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

// MessagingServiceMessagesRespondCompletedGetReader is a Reader for the MessagingServiceMessagesRespondCompletedGet structure.
type MessagingServiceMessagesRespondCompletedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MessagingServiceMessagesRespondCompletedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMessagingServiceMessagesRespondCompletedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMessagingServiceMessagesRespondCompletedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMessagingServiceMessagesRespondCompletedGetOK creates a MessagingServiceMessagesRespondCompletedGetOK with default headers values
func NewMessagingServiceMessagesRespondCompletedGetOK() *MessagingServiceMessagesRespondCompletedGetOK {
	return &MessagingServiceMessagesRespondCompletedGetOK{}
}

/*MessagingServiceMessagesRespondCompletedGetOK handles this case with default header values.

MessagingServiceMessagesRespondCompletedGetOK messaging service messages respond completed get o k
*/
type MessagingServiceMessagesRespondCompletedGetOK struct {
	Payload []*models.MessageCounter
}

func (o *MessagingServiceMessagesRespondCompletedGetOK) GetPayload() []*models.MessageCounter {
	return o.Payload
}

func (o *MessagingServiceMessagesRespondCompletedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMessagingServiceMessagesRespondCompletedGetDefault creates a MessagingServiceMessagesRespondCompletedGetDefault with default headers values
func NewMessagingServiceMessagesRespondCompletedGetDefault(code int) *MessagingServiceMessagesRespondCompletedGetDefault {
	return &MessagingServiceMessagesRespondCompletedGetDefault{
		_statusCode: code,
	}
}

/*MessagingServiceMessagesRespondCompletedGetDefault handles this case with default header values.

internal server error
*/
type MessagingServiceMessagesRespondCompletedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the messaging service messages respond completed get default response
func (o *MessagingServiceMessagesRespondCompletedGetDefault) Code() int {
	return o._statusCode
}

func (o *MessagingServiceMessagesRespondCompletedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *MessagingServiceMessagesRespondCompletedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *MessagingServiceMessagesRespondCompletedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
