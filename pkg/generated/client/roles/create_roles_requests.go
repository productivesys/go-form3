// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// Client.CreateRoles creates a new CreateRolesRequest object
// with the default values initialized.
func (c *Client) CreateRoles() *CreateRolesRequest {
	var ()
	return &CreateRolesRequest{

		RoleCreation: models.RoleCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateRolesRequest struct {

	/*RoleCreationRequest*/

	*models.RoleCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateRolesRequest) FromJson(j string) *CreateRolesRequest {

	var m models.RoleCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.RoleCreation = &m

	return o
}

func (o *CreateRolesRequest) WithRoleCreationRequest(roleCreationRequest models.RoleCreation) *CreateRolesRequest {

	o.RoleCreation = &roleCreationRequest

	return o
}

func (o *CreateRolesRequest) WithoutRoleCreationRequest() *CreateRolesRequest {

	o.RoleCreation = &models.RoleCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create roles Request
func (o *CreateRolesRequest) WithContext(ctx context.Context) *CreateRolesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create roles Request
func (o *CreateRolesRequest) WithHTTPClient(client *http.Client) *CreateRolesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateRolesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.RoleCreation != nil {
		if err := r.SetBodyParam(o.RoleCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
