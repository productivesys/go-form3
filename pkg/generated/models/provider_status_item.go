// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProviderStatusItem provider status item
// swagger:model ProviderStatusItem
type ProviderStatusItem struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// related resource id
	// Format: uuid
	RelatedResourceID strfmt.UUID `json:"related_resource_id,omitempty"`

	// related resource type
	RelatedResourceType string `json:"related_resource_type,omitempty"`
}

func ProviderStatusItemWithDefaults(defaults client.Defaults) *ProviderStatusItem {
	return &ProviderStatusItem{

		Code: defaults.GetString("ProviderStatusItem", "code"),

		Description: defaults.GetString("ProviderStatusItem", "description"),

		RelatedResourceID: defaults.GetStrfmtUUID("ProviderStatusItem", "related_resource_id"),

		RelatedResourceType: defaults.GetString("ProviderStatusItem", "related_resource_type"),
	}
}

func (m *ProviderStatusItem) WithCode(code string) *ProviderStatusItem {

	m.Code = code

	return m
}

func (m *ProviderStatusItem) WithDescription(description string) *ProviderStatusItem {

	m.Description = description

	return m
}

func (m *ProviderStatusItem) WithRelatedResourceID(relatedResourceID strfmt.UUID) *ProviderStatusItem {

	m.RelatedResourceID = relatedResourceID

	return m
}

func (m *ProviderStatusItem) WithRelatedResourceType(relatedResourceType string) *ProviderStatusItem {

	m.RelatedResourceType = relatedResourceType

	return m
}

// Validate validates this provider status item
func (m *ProviderStatusItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelatedResourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProviderStatusItem) validateRelatedResourceID(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedResourceID) { // not required
		return nil
	}

	if err := validate.FormatOf("related_resource_id", "body", "uuid", m.RelatedResourceID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProviderStatusItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderStatusItem) UnmarshalBinary(b []byte) error {
	var res ProviderStatusItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ProviderStatusItem) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
