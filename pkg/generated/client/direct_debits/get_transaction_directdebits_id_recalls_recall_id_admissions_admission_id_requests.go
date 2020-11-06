// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID creates a new GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest object
// with the default values initialized.
func (c *Client) GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID() *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {
	var ()
	return &GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID", "id"),

		RecallID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID", "recallId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest struct {

	/*AdmissionID      Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*RecallID      Recall Id      */

	RecallID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) FromJson(j string) *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {

	return o
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithAdmissionID(admissionID strfmt.UUID) *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithID(id strfmt.UUID) *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {

	o.ID = id

	return o
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithRecallID(recallID strfmt.UUID) *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {

	o.RecallID = recallID

	return o
}

//////////////////
// WithContext adds the context to the get transaction directdebits ID recalls recall ID admissions admission ID Request
func (o *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithContext(ctx context.Context) *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get transaction directdebits ID recalls recall ID admissions admission ID Request
func (o *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithHTTPClient(client *http.Client) *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param admissionId
	if err := r.SetPathParam("admissionId", o.AdmissionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}