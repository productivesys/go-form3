// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContactRelationships contact relationships
// swagger:model ContactRelationships
type ContactRelationships struct {

	// contact account
	ContactAccount *RelationshipsContactAccountRef `json:"contact_account,omitempty"`

	// party
	Party *RelationshipsPartyRef `json:"party,omitempty"`
}

func ContactRelationshipsWithDefaults(defaults client.Defaults) *ContactRelationships {
	return &ContactRelationships{

		ContactAccount: RelationshipsContactAccountRefWithDefaults(defaults),

		Party: RelationshipsPartyRefWithDefaults(defaults),
	}
}

func (m *ContactRelationships) WithContactAccount(contactAccount RelationshipsContactAccountRef) *ContactRelationships {

	m.ContactAccount = &contactAccount

	return m
}

func (m *ContactRelationships) WithoutContactAccount() *ContactRelationships {
	m.ContactAccount = nil
	return m
}

func (m *ContactRelationships) WithParty(party RelationshipsPartyRef) *ContactRelationships {

	m.Party = &party

	return m
}

func (m *ContactRelationships) WithoutParty() *ContactRelationships {
	m.Party = nil
	return m
}

// Validate validates this contact relationships
func (m *ContactRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactRelationships) validateContactAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactAccount) { // not required
		return nil
	}

	if m.ContactAccount != nil {
		if err := m.ContactAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact_account")
			}
			return err
		}
	}

	return nil
}

func (m *ContactRelationships) validateParty(formats strfmt.Registry) error {

	if swag.IsZero(m.Party) { // not required
		return nil
	}

	if m.Party != nil {
		if err := m.Party.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("party")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactRelationships) UnmarshalBinary(b []byte) error {
	var res ContactRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ContactRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}