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

// CacheServiceSaveCachesPostReader is a Reader for the CacheServiceSaveCachesPost structure.
type CacheServiceSaveCachesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceSaveCachesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceSaveCachesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceSaveCachesPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceSaveCachesPostOK creates a CacheServiceSaveCachesPostOK with default headers values
func NewCacheServiceSaveCachesPostOK() *CacheServiceSaveCachesPostOK {
	return &CacheServiceSaveCachesPostOK{}
}

/*CacheServiceSaveCachesPostOK handles this case with default header values.

CacheServiceSaveCachesPostOK cache service save caches post o k
*/
type CacheServiceSaveCachesPostOK struct {
}

func (o *CacheServiceSaveCachesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCacheServiceSaveCachesPostDefault creates a CacheServiceSaveCachesPostDefault with default headers values
func NewCacheServiceSaveCachesPostDefault(code int) *CacheServiceSaveCachesPostDefault {
	return &CacheServiceSaveCachesPostDefault{
		_statusCode: code,
	}
}

/*CacheServiceSaveCachesPostDefault handles this case with default header values.

internal server error
*/
type CacheServiceSaveCachesPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service save caches post default response
func (o *CacheServiceSaveCachesPostDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceSaveCachesPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceSaveCachesPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceSaveCachesPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
