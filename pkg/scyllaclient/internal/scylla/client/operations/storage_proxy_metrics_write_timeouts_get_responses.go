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

// StorageProxyMetricsWriteTimeoutsGetReader is a Reader for the StorageProxyMetricsWriteTimeoutsGet structure.
type StorageProxyMetricsWriteTimeoutsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsWriteTimeoutsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsWriteTimeoutsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsWriteTimeoutsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsWriteTimeoutsGetOK creates a StorageProxyMetricsWriteTimeoutsGetOK with default headers values
func NewStorageProxyMetricsWriteTimeoutsGetOK() *StorageProxyMetricsWriteTimeoutsGetOK {
	return &StorageProxyMetricsWriteTimeoutsGetOK{}
}

/*
StorageProxyMetricsWriteTimeoutsGetOK handles this case with default header values.

StorageProxyMetricsWriteTimeoutsGetOK storage proxy metrics write timeouts get o k
*/
type StorageProxyMetricsWriteTimeoutsGetOK struct {
	Payload int32
}

func (o *StorageProxyMetricsWriteTimeoutsGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMetricsWriteTimeoutsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsWriteTimeoutsGetDefault creates a StorageProxyMetricsWriteTimeoutsGetDefault with default headers values
func NewStorageProxyMetricsWriteTimeoutsGetDefault(code int) *StorageProxyMetricsWriteTimeoutsGetDefault {
	return &StorageProxyMetricsWriteTimeoutsGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyMetricsWriteTimeoutsGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsWriteTimeoutsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics write timeouts get default response
func (o *StorageProxyMetricsWriteTimeoutsGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsWriteTimeoutsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsWriteTimeoutsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsWriteTimeoutsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
