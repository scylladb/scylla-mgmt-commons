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

// FindConfigMurmur3PartitionerIgnoreMsbBitsReader is a Reader for the FindConfigMurmur3PartitionerIgnoreMsbBits structure.
type FindConfigMurmur3PartitionerIgnoreMsbBitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMurmur3PartitionerIgnoreMsbBitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMurmur3PartitionerIgnoreMsbBitsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMurmur3PartitionerIgnoreMsbBitsOK creates a FindConfigMurmur3PartitionerIgnoreMsbBitsOK with default headers values
func NewFindConfigMurmur3PartitionerIgnoreMsbBitsOK() *FindConfigMurmur3PartitionerIgnoreMsbBitsOK {
	return &FindConfigMurmur3PartitionerIgnoreMsbBitsOK{}
}

/*FindConfigMurmur3PartitionerIgnoreMsbBitsOK handles this case with default header values.

Config value
*/
type FindConfigMurmur3PartitionerIgnoreMsbBitsOK struct {
	Payload int64
}

func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMurmur3PartitionerIgnoreMsbBitsDefault creates a FindConfigMurmur3PartitionerIgnoreMsbBitsDefault with default headers values
func NewFindConfigMurmur3PartitionerIgnoreMsbBitsDefault(code int) *FindConfigMurmur3PartitionerIgnoreMsbBitsDefault {
	return &FindConfigMurmur3PartitionerIgnoreMsbBitsDefault{
		_statusCode: code,
	}
}

/*FindConfigMurmur3PartitionerIgnoreMsbBitsDefault handles this case with default header values.

unexpected error
*/
type FindConfigMurmur3PartitionerIgnoreMsbBitsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config murmur3 partitioner ignore msb bits default response
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
