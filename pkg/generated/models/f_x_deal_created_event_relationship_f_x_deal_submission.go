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

// FXDealCreatedEventRelationshipFXDealSubmission f x deal created event relationship f x deal submission
// swagger:model FXDealCreatedEventRelationshipFXDealSubmission
type FXDealCreatedEventRelationshipFXDealSubmission struct {

	// data
	// Required: true
	Data []*FXDealCreatedEventRelationshipFXDealSubmissionDataItems0 `json:"data"`
}

func FXDealCreatedEventRelationshipFXDealSubmissionWithDefaults(defaults client.Defaults) *FXDealCreatedEventRelationshipFXDealSubmission {
	return &FXDealCreatedEventRelationshipFXDealSubmission{

		Data: make([]*FXDealCreatedEventRelationshipFXDealSubmissionDataItems0, 0),
	}
}

func (m *FXDealCreatedEventRelationshipFXDealSubmission) WithData(data []*FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) *FXDealCreatedEventRelationshipFXDealSubmission {

	m.Data = data

	return m
}

// Validate validates this f x deal created event relationship f x deal submission
func (m *FXDealCreatedEventRelationshipFXDealSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXDealCreatedEventRelationshipFXDealSubmission) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
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

// MarshalBinary interface implementation
func (m *FXDealCreatedEventRelationshipFXDealSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXDealCreatedEventRelationshipFXDealSubmission) UnmarshalBinary(b []byte) error {
	var res FXDealCreatedEventRelationshipFXDealSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXDealCreatedEventRelationshipFXDealSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// FXDealCreatedEventRelationshipFXDealSubmissionDataItems0 f x deal created event relationship f x deal submission data items0
// swagger:model FXDealCreatedEventRelationshipFXDealSubmissionDataItems0
type FXDealCreatedEventRelationshipFXDealSubmissionDataItems0 struct {

	// The FX Deal Submission for this FX Deal
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type FXDealSubmissionResourceType `json:"type"`
}

func FXDealCreatedEventRelationshipFXDealSubmissionDataItems0WithDefaults(defaults client.Defaults) *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0 {
	return &FXDealCreatedEventRelationshipFXDealSubmissionDataItems0{

		ID: defaults.GetStrfmtUUIDPtr("FXDealCreatedEventRelationshipFXDealSubmissionDataItems0", "id"),

		// TODO Type: FXDealSubmissionResourceType,

	}
}

func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) WithID(id strfmt.UUID) *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0 {

	m.ID = &id

	return m
}

func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) WithoutID() *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0 {
	m.ID = nil
	return m
}

func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) WithType(typeVar FXDealSubmissionResourceType) *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0 {

	m.Type = typeVar

	return m
}

// Validate validates this f x deal created event relationship f x deal submission data items0
func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) UnmarshalBinary(b []byte) error {
	var res FXDealCreatedEventRelationshipFXDealSubmissionDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXDealCreatedEventRelationshipFXDealSubmissionDataItems0) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}