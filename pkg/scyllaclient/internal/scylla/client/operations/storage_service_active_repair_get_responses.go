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

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// StorageServiceActiveRepairGetReader is a Reader for the StorageServiceActiveRepairGet structure.
type StorageServiceActiveRepairGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceActiveRepairGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceActiveRepairGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceActiveRepairGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceActiveRepairGetOK creates a StorageServiceActiveRepairGetOK with default headers values
func NewStorageServiceActiveRepairGetOK() *StorageServiceActiveRepairGetOK {
	return &StorageServiceActiveRepairGetOK{}
}

/*
StorageServiceActiveRepairGetOK handles this case with default header values.

StorageServiceActiveRepairGetOK storage service active repair get o k
*/
type StorageServiceActiveRepairGetOK struct {
	Payload []int32
}

func (o *StorageServiceActiveRepairGetOK) GetPayload() []int32 {
	return o.Payload
}

func (o *StorageServiceActiveRepairGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceActiveRepairGetDefault creates a StorageServiceActiveRepairGetDefault with default headers values
func NewStorageServiceActiveRepairGetDefault(code int) *StorageServiceActiveRepairGetDefault {
	return &StorageServiceActiveRepairGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceActiveRepairGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceActiveRepairGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service active repair get default response
func (o *StorageServiceActiveRepairGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceActiveRepairGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceActiveRepairGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceActiveRepairGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
