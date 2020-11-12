// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package account_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new account request API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for account request API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get account requests API
*/
func (a *GetAccountRequestsRequest) Do() (*GetAccountRequestsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountRequests",
		Method:             "GET",
		PathPattern:        "/organisation/accountrequests",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAccountRequestsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountRequestsOK), nil

}

func (a *GetAccountRequestsRequest) MustDo() *GetAccountRequestsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get account requests ID API
*/
func (a *GetAccountRequestsIDRequest) Do() (*GetAccountRequestsIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountRequestsID",
		Method:             "GET",
		PathPattern:        "/organisation/accountrequests/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAccountRequestsIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountRequestsIDOK), nil

}

func (a *GetAccountRequestsIDRequest) MustDo() *GetAccountRequestsIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get account requests ID submissions submission ID API
*/
func (a *GetAccountRequestsIDSubmissionsSubmissionIDRequest) Do() (*GetAccountRequestsIDSubmissionsSubmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountRequestsIDSubmissionsSubmissionID",
		Method:             "GET",
		PathPattern:        "/organisation/accountrequests/{id}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAccountRequestsIDSubmissionsSubmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountRequestsIDSubmissionsSubmissionIDOK), nil

}

func (a *GetAccountRequestsIDSubmissionsSubmissionIDRequest) MustDo() *GetAccountRequestsIDSubmissionsSubmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post account requests API
*/
func (a *PostAccountRequestsRequest) Do() (*PostAccountRequestsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAccountRequests",
		Method:             "POST",
		PathPattern:        "/organisation/accountrequests",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostAccountRequestsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAccountRequestsCreated), nil

}

func (a *PostAccountRequestsRequest) MustDo() *PostAccountRequestsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post account requests ID submissions API
*/
func (a *PostAccountRequestsIDSubmissionsRequest) Do() (*PostAccountRequestsIDSubmissionsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAccountRequestsIDSubmissions",
		Method:             "POST",
		PathPattern:        "/organisation/accountrequests/{id}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostAccountRequestsIDSubmissionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAccountRequestsIDSubmissionsCreated), nil

}

func (a *PostAccountRequestsIDSubmissionsRequest) MustDo() *PostAccountRequestsIDSubmissionsCreated {
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
