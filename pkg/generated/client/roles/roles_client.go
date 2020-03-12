// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new roles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for roles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create roles API
*/
func (a *CreateRolesRequest) Do() (*CreateRolesCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRoles",
		Method:             "POST",
		PathPattern:        "/security/roles",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateRolesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRolesCreated), nil

}

func (a *CreateRolesRequest) MustDo() *CreateRolesCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
delete role API
*/
func (a *DeleteRoleRequest) Do() (*DeleteRoleNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRole",
		Method:             "DELETE",
		PathPattern:        "/security/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &DeleteRoleReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRoleNoContent), nil

}

func (a *DeleteRoleRequest) MustDo() *DeleteRoleNoContent {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get role API
*/
func (a *GetRoleRequest) Do() (*GetRoleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRole",
		Method:             "GET",
		PathPattern:        "/security/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetRoleReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRoleOK), nil

}

func (a *GetRoleRequest) MustDo() *GetRoleOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list roles API
*/
func (a *ListRolesRequest) Do() (*ListRolesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListRoles",
		Method:             "GET",
		PathPattern:        "/security/roles",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListRolesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRolesOK), nil

}

func (a *ListRolesRequest) MustDo() *ListRolesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
