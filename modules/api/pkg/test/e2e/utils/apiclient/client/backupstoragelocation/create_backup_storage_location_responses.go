// Code generated by go-swagger; DO NOT EDIT.

package backupstoragelocation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateBackupStorageLocationReader is a Reader for the CreateBackupStorageLocation structure.
type CreateBackupStorageLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBackupStorageLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBackupStorageLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateBackupStorageLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBackupStorageLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateBackupStorageLocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBackupStorageLocationOK creates a CreateBackupStorageLocationOK with default headers values
func NewCreateBackupStorageLocationOK() *CreateBackupStorageLocationOK {
	return &CreateBackupStorageLocationOK{}
}

/*
CreateBackupStorageLocationOK describes a response with status code 200, with default header values.

BackupStorageLocation
*/
type CreateBackupStorageLocationOK struct {
	Payload *models.BackupStorageLocation
}

// IsSuccess returns true when this create backup storage location o k response has a 2xx status code
func (o *CreateBackupStorageLocationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create backup storage location o k response has a 3xx status code
func (o *CreateBackupStorageLocationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backup storage location o k response has a 4xx status code
func (o *CreateBackupStorageLocationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create backup storage location o k response has a 5xx status code
func (o *CreateBackupStorageLocationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create backup storage location o k response a status code equal to that given
func (o *CreateBackupStorageLocationOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateBackupStorageLocationOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocationOK  %+v", 200, o.Payload)
}

func (o *CreateBackupStorageLocationOK) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocationOK  %+v", 200, o.Payload)
}

func (o *CreateBackupStorageLocationOK) GetPayload() *models.BackupStorageLocation {
	return o.Payload
}

func (o *CreateBackupStorageLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupStorageLocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBackupStorageLocationUnauthorized creates a CreateBackupStorageLocationUnauthorized with default headers values
func NewCreateBackupStorageLocationUnauthorized() *CreateBackupStorageLocationUnauthorized {
	return &CreateBackupStorageLocationUnauthorized{}
}

/*
CreateBackupStorageLocationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateBackupStorageLocationUnauthorized struct {
}

// IsSuccess returns true when this create backup storage location unauthorized response has a 2xx status code
func (o *CreateBackupStorageLocationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create backup storage location unauthorized response has a 3xx status code
func (o *CreateBackupStorageLocationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backup storage location unauthorized response has a 4xx status code
func (o *CreateBackupStorageLocationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create backup storage location unauthorized response has a 5xx status code
func (o *CreateBackupStorageLocationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create backup storage location unauthorized response a status code equal to that given
func (o *CreateBackupStorageLocationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateBackupStorageLocationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocationUnauthorized ", 401)
}

func (o *CreateBackupStorageLocationUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocationUnauthorized ", 401)
}

func (o *CreateBackupStorageLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBackupStorageLocationForbidden creates a CreateBackupStorageLocationForbidden with default headers values
func NewCreateBackupStorageLocationForbidden() *CreateBackupStorageLocationForbidden {
	return &CreateBackupStorageLocationForbidden{}
}

/*
CreateBackupStorageLocationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateBackupStorageLocationForbidden struct {
}

// IsSuccess returns true when this create backup storage location forbidden response has a 2xx status code
func (o *CreateBackupStorageLocationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create backup storage location forbidden response has a 3xx status code
func (o *CreateBackupStorageLocationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backup storage location forbidden response has a 4xx status code
func (o *CreateBackupStorageLocationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create backup storage location forbidden response has a 5xx status code
func (o *CreateBackupStorageLocationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create backup storage location forbidden response a status code equal to that given
func (o *CreateBackupStorageLocationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateBackupStorageLocationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocationForbidden ", 403)
}

func (o *CreateBackupStorageLocationForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocationForbidden ", 403)
}

func (o *CreateBackupStorageLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBackupStorageLocationDefault creates a CreateBackupStorageLocationDefault with default headers values
func NewCreateBackupStorageLocationDefault(code int) *CreateBackupStorageLocationDefault {
	return &CreateBackupStorageLocationDefault{
		_statusCode: code,
	}
}

/*
CreateBackupStorageLocationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateBackupStorageLocationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create backup storage location default response
func (o *CreateBackupStorageLocationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create backup storage location default response has a 2xx status code
func (o *CreateBackupStorageLocationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create backup storage location default response has a 3xx status code
func (o *CreateBackupStorageLocationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create backup storage location default response has a 4xx status code
func (o *CreateBackupStorageLocationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create backup storage location default response has a 5xx status code
func (o *CreateBackupStorageLocationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create backup storage location default response a status code equal to that given
func (o *CreateBackupStorageLocationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateBackupStorageLocationDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocation default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBackupStorageLocationDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/backupstoragelocation][%d] createBackupStorageLocation default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBackupStorageLocationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateBackupStorageLocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
