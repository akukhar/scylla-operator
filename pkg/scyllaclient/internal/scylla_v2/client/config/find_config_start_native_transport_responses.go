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

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigStartNativeTransportReader is a Reader for the FindConfigStartNativeTransport structure.
type FindConfigStartNativeTransportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigStartNativeTransportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigStartNativeTransportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigStartNativeTransportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigStartNativeTransportOK creates a FindConfigStartNativeTransportOK with default headers values
func NewFindConfigStartNativeTransportOK() *FindConfigStartNativeTransportOK {
	return &FindConfigStartNativeTransportOK{}
}

/*
FindConfigStartNativeTransportOK handles this case with default header values.

Config value
*/
type FindConfigStartNativeTransportOK struct {
	Payload bool
}

func (o *FindConfigStartNativeTransportOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigStartNativeTransportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigStartNativeTransportDefault creates a FindConfigStartNativeTransportDefault with default headers values
func NewFindConfigStartNativeTransportDefault(code int) *FindConfigStartNativeTransportDefault {
	return &FindConfigStartNativeTransportDefault{
		_statusCode: code,
	}
}

/*
FindConfigStartNativeTransportDefault handles this case with default header values.

unexpected error
*/
type FindConfigStartNativeTransportDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config start native transport default response
func (o *FindConfigStartNativeTransportDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigStartNativeTransportDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigStartNativeTransportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigStartNativeTransportDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
