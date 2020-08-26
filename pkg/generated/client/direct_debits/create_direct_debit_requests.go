// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

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

// Client.CreateDirectDebit creates a new CreateDirectDebitRequest object
// with the default values initialized.
func (c *Client) CreateDirectDebit() *CreateDirectDebitRequest {
	var ()
	return &CreateDirectDebitRequest{

		DirectDebitCreation: models.DirectDebitCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateDirectDebitRequest struct {

	/*DirectDebitCreationRequest*/

	*models.DirectDebitCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateDirectDebitRequest) FromJson(j string) *CreateDirectDebitRequest {

	var m models.DirectDebitCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.DirectDebitCreation = &m

	return o
}

func (o *CreateDirectDebitRequest) WithDirectDebitCreationRequest(directDebitCreationRequest models.DirectDebitCreation) *CreateDirectDebitRequest {

	o.DirectDebitCreation = &directDebitCreationRequest

	return o
}

func (o *CreateDirectDebitRequest) WithoutDirectDebitCreationRequest() *CreateDirectDebitRequest {

	o.DirectDebitCreation = &models.DirectDebitCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create direct debit Request
func (o *CreateDirectDebitRequest) WithContext(ctx context.Context) *CreateDirectDebitRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create direct debit Request
func (o *CreateDirectDebitRequest) WithHTTPClient(client *http.Client) *CreateDirectDebitRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateDirectDebitRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
