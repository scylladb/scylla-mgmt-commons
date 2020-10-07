// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewColumnFamilyMetricsReadLatencyHistogramByNameGetParams creates a new ColumnFamilyMetricsReadLatencyHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadLatencyHistogramByNameGetParams() *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadLatencyHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsReadLatencyHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadLatencyHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadLatencyHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsReadLatencyHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadLatencyHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadLatencyHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadLatencyHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadLatencyHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsReadLatencyHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics read latency histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadLatencyHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics read latency histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
