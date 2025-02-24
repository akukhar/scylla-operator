// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigExperimentalParams creates a new FindConfigExperimentalParams object
// with the default values initialized.
func NewFindConfigExperimentalParams() *FindConfigExperimentalParams {

	return &FindConfigExperimentalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigExperimentalParamsWithTimeout creates a new FindConfigExperimentalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigExperimentalParamsWithTimeout(timeout time.Duration) *FindConfigExperimentalParams {

	return &FindConfigExperimentalParams{

		timeout: timeout,
	}
}

// NewFindConfigExperimentalParamsWithContext creates a new FindConfigExperimentalParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigExperimentalParamsWithContext(ctx context.Context) *FindConfigExperimentalParams {

	return &FindConfigExperimentalParams{

		Context: ctx,
	}
}

// NewFindConfigExperimentalParamsWithHTTPClient creates a new FindConfigExperimentalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigExperimentalParamsWithHTTPClient(client *http.Client) *FindConfigExperimentalParams {

	return &FindConfigExperimentalParams{
		HTTPClient: client,
	}
}

/*
FindConfigExperimentalParams contains all the parameters to send to the API endpoint
for the find config experimental operation typically these are written to a http.Request
*/
type FindConfigExperimentalParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config experimental params
func (o *FindConfigExperimentalParams) WithTimeout(timeout time.Duration) *FindConfigExperimentalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config experimental params
func (o *FindConfigExperimentalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config experimental params
func (o *FindConfigExperimentalParams) WithContext(ctx context.Context) *FindConfigExperimentalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config experimental params
func (o *FindConfigExperimentalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config experimental params
func (o *FindConfigExperimentalParams) WithHTTPClient(client *http.Client) *FindConfigExperimentalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config experimental params
func (o *FindConfigExperimentalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigExperimentalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
