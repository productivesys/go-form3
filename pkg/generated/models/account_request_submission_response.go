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

// AccountRequestSubmissionResponse account request submission response
// swagger:model AccountRequestSubmissionResponse
type AccountRequestSubmissionResponse struct {

	// data
	Data *AccountRequestSubmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func AccountRequestSubmissionResponseWithDefaults(defaults client.Defaults) *AccountRequestSubmissionResponse {
	return &AccountRequestSubmissionResponse{

		Data: AccountRequestSubmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *AccountRequestSubmissionResponse) WithData(data AccountRequestSubmission) *AccountRequestSubmissionResponse {

	m.Data = &data

	return m
}

func (m *AccountRequestSubmissionResponse) WithoutData() *AccountRequestSubmissionResponse {
	m.Data = nil
	return m
}

func (m *AccountRequestSubmissionResponse) WithLinks(links Links) *AccountRequestSubmissionResponse {

	m.Links = &links

	return m
}

func (m *AccountRequestSubmissionResponse) WithoutLinks() *AccountRequestSubmissionResponse {
	m.Links = nil
	return m
}

// Validate validates this account request submission response
func (m *AccountRequestSubmissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRequestSubmissionResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *AccountRequestSubmissionResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequestSubmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequestSubmissionResponse) UnmarshalBinary(b []byte) error {
	var res AccountRequestSubmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRequestSubmissionResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
