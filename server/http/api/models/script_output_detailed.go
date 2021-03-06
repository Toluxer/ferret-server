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

// ScriptOutputDetailed Script Entity
// swagger:model script-output-detailed
type ScriptOutputDetailed struct {
	ScriptOutputDetailedAllOf0

	// description
	// Max Length: 255
	// Min Length: 10
	Description string `json:"description,omitempty"`

	// execution
	// Required: true
	Execution *ScriptOutputDetailedAO1Execution `json:"execution"`

	// name
	// Required: true
	// Max Length: 100
	// Min Length: 3
	Name *string `json:"name"`

	// persistence
	// Required: true
	Persistence *ScriptOutputDetailedAO1Persistence `json:"persistence"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ScriptOutputDetailed) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ScriptOutputDetailedAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ScriptOutputDetailedAllOf0 = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		Execution *ScriptOutputDetailedAO1Execution `json:"execution"`

		Name *string `json:"name"`

		Persistence *ScriptOutputDetailedAO1Persistence `json:"persistence"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.Execution = dataAO1.Execution

	m.Name = dataAO1.Name

	m.Persistence = dataAO1.Persistence

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ScriptOutputDetailed) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ScriptOutputDetailedAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		Execution *ScriptOutputDetailedAO1Execution `json:"execution"`

		Name *string `json:"name"`

		Persistence *ScriptOutputDetailedAO1Persistence `json:"persistence"`
	}

	dataAO1.Description = m.Description

	dataAO1.Execution = m.Execution

	dataAO1.Name = m.Name

	dataAO1.Persistence = m.Persistence

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this script output detailed
func (m *ScriptOutputDetailed) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ScriptOutputDetailedAllOf0
	if err := m.ScriptOutputDetailedAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptOutputDetailed) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *ScriptOutputDetailed) validateExecution(formats strfmt.Registry) error {

	if err := validate.Required("execution", "body", m.Execution); err != nil {
		return err
	}

	if m.Execution != nil {
		if err := m.Execution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execution")
			}
			return err
		}
	}

	return nil
}

func (m *ScriptOutputDetailed) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *ScriptOutputDetailed) validatePersistence(formats strfmt.Registry) error {

	if err := validate.Required("persistence", "body", m.Persistence); err != nil {
		return err
	}

	if m.Persistence != nil {
		if err := m.Persistence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("persistence")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScriptOutputDetailed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptOutputDetailed) UnmarshalBinary(b []byte) error {
	var res ScriptOutputDetailed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ScriptOutputDetailedAO1Execution Script Execution Settings
//
// Represents script execution settings like query and params
// swagger:model ScriptOutputDetailedAO1Execution
type ScriptOutputDetailedAO1Execution struct {

	// params
	Params map[string]interface{} `json:"params,omitempty"`

	// query
	// Required: true
	// Min Length: 8
	Query *string `json:"query"`
}

// Validate validates this script output detailed a o1 execution
func (m *ScriptOutputDetailedAO1Execution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptOutputDetailedAO1Execution) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("execution"+"."+"query", "body", m.Query); err != nil {
		return err
	}

	if err := validate.MinLength("execution"+"."+"query", "body", string(*m.Query), 8); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScriptOutputDetailedAO1Execution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptOutputDetailedAO1Execution) UnmarshalBinary(b []byte) error {
	var res ScriptOutputDetailedAO1Execution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ScriptOutputDetailedAO1Persistence Script Persistence
// swagger:model ScriptOutputDetailedAO1Persistence
type ScriptOutputDetailedAO1Persistence struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this script output detailed a o1 persistence
func (m *ScriptOutputDetailedAO1Persistence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptOutputDetailedAO1Persistence) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("persistence"+"."+"enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScriptOutputDetailedAO1Persistence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptOutputDetailedAO1Persistence) UnmarshalBinary(b []byte) error {
	var res ScriptOutputDetailedAO1Persistence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ScriptOutputDetailedAllOf0 Entity
//
// Represents a database entity
// swagger:model ScriptOutputDetailedAllOf0
type ScriptOutputDetailedAllOf0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// rev
	// Required: true
	Rev *string `json:"rev"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ScriptOutputDetailedAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ID = dataAO0.ID

	m.Rev = dataAO0.Rev

	// AO1
	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CreatedAt = dataAO1.CreatedAt

	m.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ScriptOutputDetailedAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}

	dataAO0.ID = m.ID

	dataAO0.Rev = m.Rev

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}

	dataAO1.CreatedAt = m.CreatedAt

	dataAO1.UpdatedAt = m.UpdatedAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this script output detailed all of0
func (m *ScriptOutputDetailedAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptOutputDetailedAllOf0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ScriptOutputDetailedAllOf0) validateRev(formats strfmt.Registry) error {

	if err := validate.Required("rev", "body", m.Rev); err != nil {
		return err
	}

	return nil
}

func (m *ScriptOutputDetailedAllOf0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScriptOutputDetailedAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptOutputDetailedAllOf0) UnmarshalBinary(b []byte) error {
	var res ScriptOutputDetailedAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
