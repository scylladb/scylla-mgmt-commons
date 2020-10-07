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

	models "github.com/scylladb/scylla-mgmt-commons/scyllaclient/agent/models"
)

// NewOperationsCheckPermissionsParams creates a new OperationsCheckPermissionsParams object
// with the default values initialized.
func NewOperationsCheckPermissionsParams() *OperationsCheckPermissionsParams {
	var ()
	return &OperationsCheckPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperationsCheckPermissionsParamsWithTimeout creates a new OperationsCheckPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperationsCheckPermissionsParamsWithTimeout(timeout time.Duration) *OperationsCheckPermissionsParams {
	var ()
	return &OperationsCheckPermissionsParams{

		timeout: timeout,
	}
}

// NewOperationsCheckPermissionsParamsWithContext creates a new OperationsCheckPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperationsCheckPermissionsParamsWithContext(ctx context.Context) *OperationsCheckPermissionsParams {
	var ()
	return &OperationsCheckPermissionsParams{

		Context: ctx,
	}
}

// NewOperationsCheckPermissionsParamsWithHTTPClient creates a new OperationsCheckPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperationsCheckPermissionsParamsWithHTTPClient(client *http.Client) *OperationsCheckPermissionsParams {
	var ()
	return &OperationsCheckPermissionsParams{
		HTTPClient: client,
	}
}

/*OperationsCheckPermissionsParams contains all the parameters to send to the API endpoint
for the operations check permissions operation typically these are written to a http.Request
*/
type OperationsCheckPermissionsParams struct {

	/*Fs
	  FS to check

	*/
	Fs *models.RemotePath

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operations check permissions params
func (o *OperationsCheckPermissionsParams) WithTimeout(timeout time.Duration) *OperationsCheckPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operations check permissions params
func (o *OperationsCheckPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operations check permissions params
func (o *OperationsCheckPermissionsParams) WithContext(ctx context.Context) *OperationsCheckPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operations check permissions params
func (o *OperationsCheckPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operations check permissions params
func (o *OperationsCheckPermissionsParams) WithHTTPClient(client *http.Client) *OperationsCheckPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operations check permissions params
func (o *OperationsCheckPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFs adds the fs to the operations check permissions params
func (o *OperationsCheckPermissionsParams) WithFs(fs *models.RemotePath) *OperationsCheckPermissionsParams {
	o.SetFs(fs)
	return o
}

// SetFs adds the fs to the operations check permissions params
func (o *OperationsCheckPermissionsParams) SetFs(fs *models.RemotePath) {
	o.Fs = fs
}

// WriteToRequest writes these params to a swagger request
func (o *OperationsCheckPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fs != nil {
		if err := r.SetBodyParam(o.Fs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
