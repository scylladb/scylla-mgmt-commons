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

// CacheServiceKeyCacheKeysToSavePostReader is a Reader for the CacheServiceKeyCacheKeysToSavePost structure.
type CacheServiceKeyCacheKeysToSavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceKeyCacheKeysToSavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceKeyCacheKeysToSavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceKeyCacheKeysToSavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceKeyCacheKeysToSavePostOK creates a CacheServiceKeyCacheKeysToSavePostOK with default headers values
func NewCacheServiceKeyCacheKeysToSavePostOK() *CacheServiceKeyCacheKeysToSavePostOK {
	return &CacheServiceKeyCacheKeysToSavePostOK{}
}

/*CacheServiceKeyCacheKeysToSavePostOK handles this case with default header values.

CacheServiceKeyCacheKeysToSavePostOK cache service key cache keys to save post o k
*/
type CacheServiceKeyCacheKeysToSavePostOK struct {
}

func (o *CacheServiceKeyCacheKeysToSavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCacheServiceKeyCacheKeysToSavePostDefault creates a CacheServiceKeyCacheKeysToSavePostDefault with default headers values
func NewCacheServiceKeyCacheKeysToSavePostDefault(code int) *CacheServiceKeyCacheKeysToSavePostDefault {
	return &CacheServiceKeyCacheKeysToSavePostDefault{
		_statusCode: code,
	}
}

/*CacheServiceKeyCacheKeysToSavePostDefault handles this case with default header values.

internal server error
*/
type CacheServiceKeyCacheKeysToSavePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service key cache keys to save post default response
func (o *CacheServiceKeyCacheKeysToSavePostDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceKeyCacheKeysToSavePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceKeyCacheKeysToSavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceKeyCacheKeysToSavePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
