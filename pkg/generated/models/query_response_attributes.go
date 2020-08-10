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
)

// QueryResponseAttributes query response attributes
// swagger:model QueryResponseAttributes
type QueryResponseAttributes struct {

	// answer
	// Required: true
	Answer QueryResponseAnswer `json:"answer"`

	// compensation amount
	CompensationAmount string `json:"compensation_amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

func QueryResponseAttributesWithDefaults(defaults client.Defaults) *QueryResponseAttributes {
	return &QueryResponseAttributes{

		// TODO Answer: QueryResponseAnswer,

		CompensationAmount: defaults.GetString("QueryResponseAttributes", "compensation_amount"),

		Currency: defaults.GetString("QueryResponseAttributes", "currency"),
	}
}

func (m *QueryResponseAttributes) WithAnswer(answer QueryResponseAnswer) *QueryResponseAttributes {

	m.Answer = answer

	return m
}

func (m *QueryResponseAttributes) WithCompensationAmount(compensationAmount string) *QueryResponseAttributes {

	m.CompensationAmount = compensationAmount

	return m
}

func (m *QueryResponseAttributes) WithCurrency(currency string) *QueryResponseAttributes {

	m.Currency = currency

	return m
}

// Validate validates this query response attributes
func (m *QueryResponseAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseAttributes) validateAnswer(formats strfmt.Registry) error {

	if err := m.Answer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("answer")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseAttributes) UnmarshalBinary(b []byte) error {
	var res QueryResponseAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}