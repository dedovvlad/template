// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTestParams creates a new TestParams object
// with the default values initialized.
func NewTestParams() *TestParams {

	return &TestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestParamsWithTimeout creates a new TestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestParamsWithTimeout(timeout time.Duration) *TestParams {

	return &TestParams{

		timeout: timeout,
	}
}

// NewTestParamsWithContext creates a new TestParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestParamsWithContext(ctx context.Context) *TestParams {

	params := TestParams{

		Context: ctx,
	}

	meshMap := ctx.Value("mesh").(map[string]string)
	if meshMap == nil {
		return &params
	}

	if url, ok := meshMap[strings.ToLower("Test")]; ok {
		params.MeshURL = url
	}

	return &params
}

// NewTestParamsWithHTTPClient creates a new TestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestParamsWithHTTPClient(client *http.Client) *TestParams {

	return &TestParams{
		HTTPClient: client,
	}
}

/*
TestParams contains all the parameters to send to the API endpoint
for the test operation typically these are written to a http.Request
*/
type TestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	Transport *http.Transport

	MeshURL string
}

// WithTimeout adds the timeout to the test params
func (o *TestParams) WithTimeout(timeout time.Duration) *TestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test params
func (o *TestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test params
func (o *TestParams) WithContext(ctx context.Context) *TestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test params
func (o *TestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test params
func (o *TestParams) WithHTTPClient(client *http.Client) *TestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test params
func (o *TestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	if o.Context != nil {
		if s, ok := tracer.SpanFromContext(o.Context); ok {
			_ = tracer.Inject(s.Context(), tracer.HTTPHeadersCarrier(r.GetHeaderParams()))
		}
	}

	return nil
}