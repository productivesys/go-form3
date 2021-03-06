// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package participants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new participants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for participants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get openbanking participants API
*/
func (a *GetOpenbankingParticipantsRequest) Do() (*GetOpenbankingParticipantsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOpenbankingParticipants",
		Method:             "GET",
		PathPattern:        "/openbanking/participants",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetOpenbankingParticipantsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOpenbankingParticipantsOK), nil

}

func (a *GetOpenbankingParticipantsRequest) MustDo() *GetOpenbankingParticipantsOK {
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
