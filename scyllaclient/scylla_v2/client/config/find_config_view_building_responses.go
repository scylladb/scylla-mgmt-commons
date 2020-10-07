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

// FindConfigViewBuildingReader is a Reader for the FindConfigViewBuilding structure.
type FindConfigViewBuildingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigViewBuildingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigViewBuildingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigViewBuildingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigViewBuildingOK creates a FindConfigViewBuildingOK with default headers values
func NewFindConfigViewBuildingOK() *FindConfigViewBuildingOK {
	return &FindConfigViewBuildingOK{}
}

/*FindConfigViewBuildingOK handles this case with default header values.

Config value
*/
type FindConfigViewBuildingOK struct {
	Payload bool
}

func (o *FindConfigViewBuildingOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigViewBuildingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigViewBuildingDefault creates a FindConfigViewBuildingDefault with default headers values
func NewFindConfigViewBuildingDefault(code int) *FindConfigViewBuildingDefault {
	return &FindConfigViewBuildingDefault{
		_statusCode: code,
	}
}

/*FindConfigViewBuildingDefault handles this case with default header values.

unexpected error
*/
type FindConfigViewBuildingDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config view building default response
func (o *FindConfigViewBuildingDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigViewBuildingDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigViewBuildingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigViewBuildingDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
