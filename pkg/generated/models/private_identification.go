// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateIdentification private identification
// swagger:model PrivateIdentification
type PrivateIdentification struct {

	// The street name and house number of the postal address of the account holder.
	Address []string `json:"address"`

	// The birth country of the account holder. Must be ISO 3166-1 code.
	// Pattern: ^[A-Z]{2}$
	BirthCountry string `json:"birth_country,omitempty"`

	// The birth date of the account holder.
	// Format: date
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`

	// The city where the postal address of the account holder is located.
	// Max Length: 35
	// Min Length: 1
	City string `json:"city,omitempty"`

	// The country where the postal address of the account holder is located. Must be ISO 3166-1 code.
	// Pattern: ^[A-Z]{2}$
	Country string `json:"country,omitempty"`

	// identification
	// Required: true
	// Max Length: 140
	// Min Length: 1
	Identification string `json:"identification"`

	// Issuer of the `identification`
	IdentificationIssuer string `json:"identification_issuer,omitempty"`

	// Scheme of the `identification`
	IdentificationScheme string `json:"identification_scheme,omitempty"`

	// The code that specifies the type of `identification`
	IdentificationSchemeCode string `json:"identification_scheme_code,omitempty"`
}

func PrivateIdentificationWithDefaults(defaults client.Defaults) *PrivateIdentification {
	return &PrivateIdentification{

		Address: make([]string, 0),

		BirthCountry: defaults.GetString("PrivateIdentification", "birth_country"),

		BirthDate: defaults.GetStrfmtDatePtr("PrivateIdentification", "birth_date"),

		City: defaults.GetString("PrivateIdentification", "city"),

		Country: defaults.GetString("PrivateIdentification", "country"),

		Identification: defaults.GetString("PrivateIdentification", "identification"),

		IdentificationIssuer: defaults.GetString("PrivateIdentification", "identification_issuer"),

		IdentificationScheme: defaults.GetString("PrivateIdentification", "identification_scheme"),

		IdentificationSchemeCode: defaults.GetString("PrivateIdentification", "identification_scheme_code"),
	}
}

func (m *PrivateIdentification) WithAddress(address []string) *PrivateIdentification {

	m.Address = address

	return m
}

func (m *PrivateIdentification) WithBirthCountry(birthCountry string) *PrivateIdentification {

	m.BirthCountry = birthCountry

	return m
}

func (m *PrivateIdentification) WithBirthDate(birthDate strfmt.Date) *PrivateIdentification {

	m.BirthDate = &birthDate

	return m
}

func (m *PrivateIdentification) WithoutBirthDate() *PrivateIdentification {
	m.BirthDate = nil
	return m
}

func (m *PrivateIdentification) WithCity(city string) *PrivateIdentification {

	m.City = city

	return m
}

func (m *PrivateIdentification) WithCountry(country string) *PrivateIdentification {

	m.Country = country

	return m
}

func (m *PrivateIdentification) WithIdentification(identification string) *PrivateIdentification {

	m.Identification = identification

	return m
}

func (m *PrivateIdentification) WithIdentificationIssuer(identificationIssuer string) *PrivateIdentification {

	m.IdentificationIssuer = identificationIssuer

	return m
}

func (m *PrivateIdentification) WithIdentificationScheme(identificationScheme string) *PrivateIdentification {

	m.IdentificationScheme = identificationScheme

	return m
}

func (m *PrivateIdentification) WithIdentificationSchemeCode(identificationSchemeCode string) *PrivateIdentification {

	m.IdentificationSchemeCode = identificationSchemeCode

	return m
}

// Validate validates this private identification
func (m *PrivateIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateIdentification) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	for i := 0; i < len(m.Address); i++ {

		if err := validate.MinLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *PrivateIdentification) validateBirthCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthCountry) { // not required
		return nil
	}

	if err := validate.Pattern("birth_country", "body", string(m.BirthCountry), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *PrivateIdentification) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PrivateIdentification) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MinLength("city", "body", string(m.City), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(m.City), 35); err != nil {
		return err
	}

	return nil
}

func (m *PrivateIdentification) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.Pattern("country", "body", string(m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *PrivateIdentification) validateIdentification(formats strfmt.Registry) error {

	if err := validate.RequiredString("identification", "body", string(m.Identification)); err != nil {
		return err
	}

	if err := validate.MinLength("identification", "body", string(m.Identification), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("identification", "body", string(m.Identification), 140); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateIdentification) UnmarshalBinary(b []byte) error {
	var res PrivateIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PrivateIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
