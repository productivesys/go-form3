// Code generated by go-swagger; DO NOT EDIT.

package claims

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.CreateClaimReversal creates a new CreateClaimReversalRequest object
// with the default values initialized.
func (c *Client) CreateClaimReversal() *CreateClaimReversalRequest {
	var ()
	return &CreateClaimReversalRequest{

		ClaimReversalCreation: models.ClaimReversalCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateClaimReversal", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateClaimReversalRequest struct {

	/*ClaimReversalCreationRequest*/

	*models.ClaimReversalCreation

	/*ID      Claim Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateClaimReversalRequest) FromJson(j string) *CreateClaimReversalRequest {

	var m models.ClaimReversalCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.ClaimReversalCreation = &m

	return o
}

func (o *CreateClaimReversalRequest) WithClaimReversalCreationRequest(claimReversalCreationRequest models.ClaimReversalCreation) *CreateClaimReversalRequest {

	o.ClaimReversalCreation = &claimReversalCreationRequest

	return o
}

func (o *CreateClaimReversalRequest) WithoutClaimReversalCreationRequest() *CreateClaimReversalRequest {

	o.ClaimReversalCreation = &models.ClaimReversalCreation{}

	return o
}

func (o *CreateClaimReversalRequest) WithID(id strfmt.UUID) *CreateClaimReversalRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create claim reversal Request
func (o *CreateClaimReversalRequest) WithContext(ctx context.Context) *CreateClaimReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create claim reversal Request
func (o *CreateClaimReversalRequest) WithHTTPClient(client *http.Client) *CreateClaimReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateClaimReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ClaimReversalCreation != nil {
		if err := r.SetBodyParam(o.ClaimReversalCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
