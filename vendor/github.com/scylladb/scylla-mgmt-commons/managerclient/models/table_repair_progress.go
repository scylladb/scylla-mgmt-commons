// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TableRepairProgress table repair progress
// swagger:model TableRepairProgress
type TableRepairProgress struct {

	// completed at
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	// error
	Error int64 `json:"error,omitempty"`

	// keyspace
	Keyspace string `json:"keyspace,omitempty"`

	// started at
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	// success
	Success int64 `json:"success,omitempty"`

	// table
	Table string `json:"table,omitempty"`

	// token ranges
	TokenRanges int64 `json:"token_ranges,omitempty"`
}

// Validate validates this table repair progress
func (m *TableRepairProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableRepairProgress) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TableRepairProgress) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TableRepairProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableRepairProgress) UnmarshalBinary(b []byte) error {
	var res TableRepairProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
