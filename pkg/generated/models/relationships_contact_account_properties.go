// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelationshipsContactAccountProperties relationships contact account properties
// swagger:model RelationshipsContactAccountProperties
type RelationshipsContactAccountProperties struct {

	// data
	Data []*ContactAccount `json:"data"`

	// The contact account which initiated this funding request
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type ContactAccountResourceType `json:"type"`
}

func RelationshipsContactAccountPropertiesWithDefaults(defaults client.Defaults) *RelationshipsContactAccountProperties {
	return &RelationshipsContactAccountProperties{

		Data: make([]*ContactAccount, 0),

		ID: defaults.GetStrfmtUUIDPtr("RelationshipsContactAccountProperties", "id"),

		// TODO Type: ContactAccountResourceType,

	}
}

func (m *RelationshipsContactAccountProperties) WithData(data []*ContactAccount) *RelationshipsContactAccountProperties {

	m.Data = data

	return m
}

func (m *RelationshipsContactAccountProperties) WithID(id strfmt.UUID) *RelationshipsContactAccountProperties {

	m.ID = &id

	return m
}

func (m *RelationshipsContactAccountProperties) WithoutID() *RelationshipsContactAccountProperties {
	m.ID = nil
	return m
}

func (m *RelationshipsContactAccountProperties) WithType(typeVar ContactAccountResourceType) *RelationshipsContactAccountProperties {

	m.Type = typeVar

	return m
}

// Validate validates this relationships contact account properties
func (m *RelationshipsContactAccountProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipsContactAccountProperties) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RelationshipsContactAccountProperties) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationshipsContactAccountProperties) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipsContactAccountProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsContactAccountProperties) UnmarshalBinary(b []byte) error {
	var res RelationshipsContactAccountProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsContactAccountProperties) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
