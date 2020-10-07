// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla_v2/models"
)

// FindConfigAuthorizerReader is a Reader for the FindConfigAuthorizer structure.
type FindConfigAuthorizerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAuthorizerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAuthorizerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAuthorizerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAuthorizerOK creates a FindConfigAuthorizerOK with default headers values
func NewFindConfigAuthorizerOK() *FindConfigAuthorizerOK {
	return &FindConfigAuthorizerOK{}
}

/*FindConfigAuthorizerOK handles this case with default header values.

Config value
*/
type FindConfigAuthorizerOK struct {
	Payload string
}

func (o *FindConfigAuthorizerOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigAuthorizerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAuthorizerDefault creates a FindConfigAuthorizerDefault with default headers values
func NewFindConfigAuthorizerDefault(code int) *FindConfigAuthorizerDefault {
	return &FindConfigAuthorizerDefault{
		_statusCode: code,
	}
}

/*FindConfigAuthorizerDefault handles this case with default header values.

unexpected error
*/
type FindConfigAuthorizerDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config authorizer default response
func (o *FindConfigAuthorizerDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAuthorizerDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAuthorizerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAuthorizerDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
