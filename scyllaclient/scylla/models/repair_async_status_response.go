// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RepairAsyncStatusResponse repair_async_statusResponse
// swagger:model repair_async_statusResponse
type RepairAsyncStatusResponse string

const (

	// RepairAsyncStatusResponseRUNNING captures enum value "RUNNING"
	RepairAsyncStatusResponseRUNNING RepairAsyncStatusResponse = "RUNNING"

	// RepairAsyncStatusResponseSUCCESSFUL captures enum value "SUCCESSFUL"
	RepairAsyncStatusResponseSUCCESSFUL RepairAsyncStatusResponse = "SUCCESSFUL"

	// RepairAsyncStatusResponseFAILED captures enum value "FAILED"
	RepairAsyncStatusResponseFAILED RepairAsyncStatusResponse = "FAILED"
)

// for schema
var repairAsyncStatusResponseEnum []interface{}

func init() {
	var res []RepairAsyncStatusResponse
	if err := json.Unmarshal([]byte(`["RUNNING","SUCCESSFUL","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		repairAsyncStatusResponseEnum = append(repairAsyncStatusResponseEnum, v)
	}
}

func (m RepairAsyncStatusResponse) validateRepairAsyncStatusResponseEnum(path, location string, value RepairAsyncStatusResponse) error {
	if err := validate.Enum(path, location, value, repairAsyncStatusResponseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this repair async status response
func (m RepairAsyncStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRepairAsyncStatusResponseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
