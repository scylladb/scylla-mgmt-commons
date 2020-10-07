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

// GetClusterClusterIDTaskBackupTaskIDRunIDReader is a Reader for the GetClusterClusterIDTaskBackupTaskIDRunID structure.
type GetClusterClusterIDTaskBackupTaskIDRunIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDTaskBackupTaskIDRunIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDTaskBackupTaskIDRunIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterClusterIDTaskBackupTaskIDRunIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDTaskBackupTaskIDRunIDOK creates a GetClusterClusterIDTaskBackupTaskIDRunIDOK with default headers values
func NewGetClusterClusterIDTaskBackupTaskIDRunIDOK() *GetClusterClusterIDTaskBackupTaskIDRunIDOK {
	return &GetClusterClusterIDTaskBackupTaskIDRunIDOK{}
}

/*GetClusterClusterIDTaskBackupTaskIDRunIDOK handles this case with default header values.

Backup progress
*/
type GetClusterClusterIDTaskBackupTaskIDRunIDOK struct {
	Payload *models.TaskRunBackupProgress
}

func (o *GetClusterClusterIDTaskBackupTaskIDRunIDOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/backup/{task_id}/{run_id}][%d] getClusterClusterIdTaskBackupTaskIdRunIdOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDTaskBackupTaskIDRunIDOK) GetPayload() *models.TaskRunBackupProgress {
	return o.Payload
}

func (o *GetClusterClusterIDTaskBackupTaskIDRunIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskRunBackupProgress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDTaskBackupTaskIDRunIDDefault creates a GetClusterClusterIDTaskBackupTaskIDRunIDDefault with default headers values
func NewGetClusterClusterIDTaskBackupTaskIDRunIDDefault(code int) *GetClusterClusterIDTaskBackupTaskIDRunIDDefault {
	return &GetClusterClusterIDTaskBackupTaskIDRunIDDefault{
		_statusCode: code,
	}
}

/*GetClusterClusterIDTaskBackupTaskIDRunIDDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDTaskBackupTaskIDRunIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID task backup task ID run ID default response
func (o *GetClusterClusterIDTaskBackupTaskIDRunIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDTaskBackupTaskIDRunIDDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/backup/{task_id}/{run_id}][%d] GetClusterClusterIDTaskBackupTaskIDRunID default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDTaskBackupTaskIDRunIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDTaskBackupTaskIDRunIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
