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

// FindConfigLsaReclamationStepReader is a Reader for the FindConfigLsaReclamationStep structure.
type FindConfigLsaReclamationStepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigLsaReclamationStepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigLsaReclamationStepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigLsaReclamationStepDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigLsaReclamationStepOK creates a FindConfigLsaReclamationStepOK with default headers values
func NewFindConfigLsaReclamationStepOK() *FindConfigLsaReclamationStepOK {
	return &FindConfigLsaReclamationStepOK{}
}

/*
FindConfigLsaReclamationStepOK handles this case with default header values.

Config value
*/
type FindConfigLsaReclamationStepOK struct {
	Payload int64
}

func (o *FindConfigLsaReclamationStepOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigLsaReclamationStepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigLsaReclamationStepDefault creates a FindConfigLsaReclamationStepDefault with default headers values
func NewFindConfigLsaReclamationStepDefault(code int) *FindConfigLsaReclamationStepDefault {
	return &FindConfigLsaReclamationStepDefault{
		_statusCode: code,
	}
}

/*
FindConfigLsaReclamationStepDefault handles this case with default header values.

unexpected error
*/
type FindConfigLsaReclamationStepDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config lsa reclamation step default response
func (o *FindConfigLsaReclamationStepDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigLsaReclamationStepDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigLsaReclamationStepDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigLsaReclamationStepDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
