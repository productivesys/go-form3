// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceType resource type
// swagger:model ResourceType
type ResourceType string

const (

	// ResourceTypeFxDeals captures enum value "fx_deals"
	ResourceTypeFxDeals ResourceType = "fx_deals"

	// ResourceTypeRecallDecisionAdmissions captures enum value "recall_decision_admissions"
	ResourceTypeRecallDecisionAdmissions ResourceType = "recall_decision_admissions"

	// ResourceTypeLimits captures enum value "limits"
	ResourceTypeLimits ResourceType = "limits"

	// ResourceTypePartyProductEvents captures enum value "party_product_events"
	ResourceTypePartyProductEvents ResourceType = "party_product_events"

	// ResourceTypePaymentDefaults captures enum value "payment_defaults"
	ResourceTypePaymentDefaults ResourceType = "payment_defaults"

	// ResourceTypeAccountRoutings captures enum value "account_routings"
	ResourceTypeAccountRoutings ResourceType = "account_routings"

	// ResourceTypePartyAccounts captures enum value "party_accounts"
	ResourceTypePartyAccounts ResourceType = "party_accounts"

	// ResourceTypeReversalSubmissionValidations captures enum value "reversal_submission_validations"
	ResourceTypeReversalSubmissionValidations ResourceType = "reversal_submission_validations"

	// ResourceTypeAccounts captures enum value "accounts"
	ResourceTypeAccounts ResourceType = "accounts"

	// ResourceTypeReturnReversalAdmissions captures enum value "return_reversal_admissions"
	ResourceTypeReturnReversalAdmissions ResourceType = "return_reversal_admissions"

	// ResourceTypePartyActors captures enum value "party_actors"
	ResourceTypePartyActors ResourceType = "party_actors"

	// ResourceTypeRecalls captures enum value "recalls"
	ResourceTypeRecalls ResourceType = "recalls"

	// ResourceTypeRecallReversalAdmissions captures enum value "recall_reversal_admissions"
	ResourceTypeRecallReversalAdmissions ResourceType = "recall_reversal_admissions"

	// ResourceTypeProductEvents captures enum value "product_events"
	ResourceTypeProductEvents ResourceType = "product_events"

	// ResourceTypeContacts captures enum value "contacts"
	ResourceTypeContacts ResourceType = "contacts"

	// ResourceTypeReturnReversals captures enum value "return_reversals"
	ResourceTypeReturnReversals ResourceType = "return_reversals"

	// ResourceTypeReversalSubmissions captures enum value "reversal_submissions"
	ResourceTypeReversalSubmissions ResourceType = "reversal_submissions"

	// ResourceTypeDirectDebit captures enum value "direct_debit"
	ResourceTypeDirectDebit ResourceType = "direct_debit"

	// ResourceTypeReturnSubmissions captures enum value "return_submissions"
	ResourceTypeReturnSubmissions ResourceType = "return_submissions"

	// ResourceTypeDirectAccount captures enum value "direct_account"
	ResourceTypeDirectAccount ResourceType = "direct_account"

	// ResourceTypeProductAssociations captures enum value "product_associations"
	ResourceTypeProductAssociations ResourceType = "product_associations"

	// ResourceTypeBankIds captures enum value "bank_ids"
	ResourceTypeBankIds ResourceType = "bank_ids"

	// ResourceTypeReturns captures enum value "returns"
	ResourceTypeReturns ResourceType = "returns"

	// ResourceTypeReturnAdmissions captures enum value "return_admissions"
	ResourceTypeReturnAdmissions ResourceType = "return_admissions"

	// ResourceTypePaymentAutomaticReturns captures enum value "payment_automatic_returns"
	ResourceTypePaymentAutomaticReturns ResourceType = "payment_automatic_returns"

	// ResourceTypeReversals captures enum value "reversals"
	ResourceTypeReversals ResourceType = "reversals"

	// ResourceTypeRecallAdmissions captures enum value "recall_admissions"
	ResourceTypeRecallAdmissions ResourceType = "recall_admissions"

	// ResourceTypeRecallSubmissionValidations captures enum value "recall_submission_validations"
	ResourceTypeRecallSubmissionValidations ResourceType = "recall_submission_validations"

	// ResourceTypeReturnSubmissionValidations captures enum value "return_submission_validations"
	ResourceTypeReturnSubmissionValidations ResourceType = "return_submission_validations"

	// ResourceTypeRecallDecisionSubmissionValidations captures enum value "recall_decision_submission_validations"
	ResourceTypeRecallDecisionSubmissionValidations ResourceType = "recall_decision_submission_validations"

	// ResourceTypePaymentAdmissions captures enum value "payment_admissions"
	ResourceTypePaymentAdmissions ResourceType = "payment_admissions"

	// ResourceTypeRecallReversals captures enum value "recall_reversals"
	ResourceTypeRecallReversals ResourceType = "recall_reversals"

	// ResourceTypePaymentAdvices captures enum value "payment_advices"
	ResourceTypePaymentAdvices ResourceType = "payment_advices"

	// ResourceTypeRecallDecisions captures enum value "recall_decisions"
	ResourceTypeRecallDecisions ResourceType = "recall_decisions"

	// ResourceTypePaymentBatches captures enum value "payment_batches"
	ResourceTypePaymentBatches ResourceType = "payment_batches"

	// ResourceTypeBics captures enum value "bics"
	ResourceTypeBics ResourceType = "bics"

	// ResourceTypePartyProducts captures enum value "party_products"
	ResourceTypePartyProducts ResourceType = "party_products"

	// ResourceTypeAccountConfigurations captures enum value "account_configurations"
	ResourceTypeAccountConfigurations ResourceType = "account_configurations"

	// ResourceTypePaymentSubmissions captures enum value "payment_submissions"
	ResourceTypePaymentSubmissions ResourceType = "payment_submissions"

	// ResourceTypeRecallSubmissions captures enum value "recall_submissions"
	ResourceTypeRecallSubmissions ResourceType = "recall_submissions"

	// ResourceTypePaymentAdviceSubmissionValidations captures enum value "payment_advice_submission_validations"
	ResourceTypePaymentAdviceSubmissionValidations ResourceType = "payment_advice_submission_validations"

	// ResourceTypeAccountEvents captures enum value "account_events"
	ResourceTypeAccountEvents ResourceType = "account_events"

	// ResourceTypePositions captures enum value "positions"
	ResourceTypePositions ResourceType = "positions"

	// ResourceTypeRecallDecisionSubmissions captures enum value "recall_decision_submissions"
	ResourceTypeRecallDecisionSubmissions ResourceType = "recall_decision_submissions"

	// ResourceTypePaymentAdviceSubmissions captures enum value "payment_advice_submissions"
	ResourceTypePaymentAdviceSubmissions ResourceType = "payment_advice_submissions"

	// ResourceTypePaymentSubmissionValidations captures enum value "payment_submission_validations"
	ResourceTypePaymentSubmissionValidations ResourceType = "payment_submission_validations"

	// ResourceTypePayments captures enum value "payments"
	ResourceTypePayments ResourceType = "payments"

	// ResourceTypeContactAccounts captures enum value "contact_accounts"
	ResourceTypeContactAccounts ResourceType = "contact_accounts"

	// ResourceTypeSchemeMessageAdmissions captures enum value "scheme_message_admissions"
	ResourceTypeSchemeMessageAdmissions ResourceType = "scheme_message_admissions"

	// ResourceTypeReversalAdmissions captures enum value "reversal_admissions"
	ResourceTypeReversalAdmissions ResourceType = "reversal_admissions"

	// ResourceTypeParties captures enum value "parties"
	ResourceTypeParties ResourceType = "parties"

	// ResourceTypeSchemeMessages captures enum value "scheme_messages"
	ResourceTypeSchemeMessages ResourceType = "scheme_messages"
)

// for schema
var resourceTypeEnum []interface{}

func init() {
	var res []ResourceType
	if err := json.Unmarshal([]byte(`["fx_deals","recall_decision_admissions","limits","party_product_events","payment_defaults","account_routings","party_accounts","reversal_submission_validations","accounts","return_reversal_admissions","party_actors","recalls","recall_reversal_admissions","product_events","contacts","return_reversals","reversal_submissions","direct_debit","return_submissions","direct_account","product_associations","bank_ids","returns","return_admissions","payment_automatic_returns","reversals","recall_admissions","recall_submission_validations","return_submission_validations","recall_decision_submission_validations","payment_admissions","recall_reversals","payment_advices","recall_decisions","payment_batches","bics","party_products","account_configurations","payment_submissions","recall_submissions","payment_advice_submission_validations","account_events","positions","recall_decision_submissions","payment_advice_submissions","payment_submission_validations","payments","contact_accounts","scheme_message_admissions","reversal_admissions","parties","scheme_messages"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeEnum = append(resourceTypeEnum, v)
	}
}

func (m ResourceType) validateResourceTypeEnum(path, location string, value ResourceType) error {
	if err := validate.Enum(path, location, value, resourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this resource type
func (m ResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
