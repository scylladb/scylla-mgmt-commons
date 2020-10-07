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

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams() *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParamsWithTimeout creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParamsWithContext creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParamsWithHTTPClient creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams contains all the parameters to send to the API endpoint
for the column family metrics write latency moving average histogram get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics write latency moving average histogram get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics write latency moving average histogram get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics write latency moving average histogram get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics write latency moving average histogram get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics write latency moving average histogram get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics write latency moving average histogram get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
