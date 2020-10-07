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

// NewStorageServiceLoadMapGetParams creates a new StorageServiceLoadMapGetParams object
// with the default values initialized.
func NewStorageServiceLoadMapGetParams() *StorageServiceLoadMapGetParams {

	return &StorageServiceLoadMapGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceLoadMapGetParamsWithTimeout creates a new StorageServiceLoadMapGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceLoadMapGetParamsWithTimeout(timeout time.Duration) *StorageServiceLoadMapGetParams {

	return &StorageServiceLoadMapGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceLoadMapGetParamsWithContext creates a new StorageServiceLoadMapGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceLoadMapGetParamsWithContext(ctx context.Context) *StorageServiceLoadMapGetParams {

	return &StorageServiceLoadMapGetParams{

		Context: ctx,
	}
}

// NewStorageServiceLoadMapGetParamsWithHTTPClient creates a new StorageServiceLoadMapGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceLoadMapGetParamsWithHTTPClient(client *http.Client) *StorageServiceLoadMapGetParams {

	return &StorageServiceLoadMapGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceLoadMapGetParams contains all the parameters to send to the API endpoint
for the storage service load map get operation typically these are written to a http.Request
*/
type StorageServiceLoadMapGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service load map get params
func (o *StorageServiceLoadMapGetParams) WithTimeout(timeout time.Duration) *StorageServiceLoadMapGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service load map get params
func (o *StorageServiceLoadMapGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service load map get params
func (o *StorageServiceLoadMapGetParams) WithContext(ctx context.Context) *StorageServiceLoadMapGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service load map get params
func (o *StorageServiceLoadMapGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service load map get params
func (o *StorageServiceLoadMapGetParams) WithHTTPClient(client *http.Client) *StorageServiceLoadMapGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service load map get params
func (o *StorageServiceLoadMapGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceLoadMapGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
