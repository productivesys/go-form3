// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new addressbook API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for addressbook API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get addressbook parties API
*/
func (a *GetAddressbookPartiesRequest) Do() (*GetAddressbookPartiesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookParties",
		Method:             "GET",
		PathPattern:        "/addressbook/parties",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesOK), nil

}

func (a *GetAddressbookPartiesRequest) MustDo() *GetAddressbookPartiesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID API
*/
func (a *GetAddressbookPartiesIDRequest) Do() (*GetAddressbookPartiesIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesID",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDOK), nil

}

func (a *GetAddressbookPartiesIDRequest) MustDo() *GetAddressbookPartiesIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID accounts party account ID API
*/
func (a *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) Do() (*GetAddressbookPartiesIDAccountsPartyAccountIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDAccountsPartyAccountID",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/accounts/{party_account_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDAccountsPartyAccountIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDAccountsPartyAccountIDOK), nil

}

func (a *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) MustDo() *GetAddressbookPartiesIDAccountsPartyAccountIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID actors API
*/
func (a *GetAddressbookPartiesIDActorsRequest) Do() (*GetAddressbookPartiesIDActorsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDActors",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/actors",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDActorsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDActorsOK), nil

}

func (a *GetAddressbookPartiesIDActorsRequest) MustDo() *GetAddressbookPartiesIDActorsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID actors party actor ID API
*/
func (a *GetAddressbookPartiesIDActorsPartyActorIDRequest) Do() (*GetAddressbookPartiesIDActorsPartyActorIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDActorsPartyActorID",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/actors/{party_actor_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDActorsPartyActorIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDActorsPartyActorIDOK), nil

}

func (a *GetAddressbookPartiesIDActorsPartyActorIDRequest) MustDo() *GetAddressbookPartiesIDActorsPartyActorIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID contacts API
*/
func (a *GetAddressbookPartiesIDContactsRequest) Do() (*GetAddressbookPartiesIDContactsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDContacts",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/contacts",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDContactsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDContactsOK), nil

}

func (a *GetAddressbookPartiesIDContactsRequest) MustDo() *GetAddressbookPartiesIDContactsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID contacts contact ID API
*/
func (a *GetAddressbookPartiesIDContactsContactIDRequest) Do() (*GetAddressbookPartiesIDContactsContactIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDContactsContactID",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/contacts/{contact_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDContactsContactIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDContactsContactIDOK), nil

}

func (a *GetAddressbookPartiesIDContactsContactIDRequest) MustDo() *GetAddressbookPartiesIDContactsContactIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID contacts contact ID accounts API
*/
func (a *GetAddressbookPartiesIDContactsContactIDAccountsRequest) Do() (*GetAddressbookPartiesIDContactsContactIDAccountsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDContactsContactIDAccounts",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/contacts/{contact_id}/accounts",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDContactsContactIDAccountsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDContactsContactIDAccountsOK), nil

}

func (a *GetAddressbookPartiesIDContactsContactIDAccountsRequest) MustDo() *GetAddressbookPartiesIDContactsContactIDAccountsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID contacts contact ID accounts contact account ID API
*/
func (a *GetAddressbookPartiesIDContactsContactIDAccountsContactAccountIDRequest) Do() (*GetAddressbookPartiesIDContactsContactIDAccountsContactAccountIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDContactsContactIDAccountsContactAccountID",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/contacts/{contact_id}/accounts/{contact_account_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDContactsContactIDAccountsContactAccountIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDContactsContactIDAccountsContactAccountIDOK), nil

}

func (a *GetAddressbookPartiesIDContactsContactIDAccountsContactAccountIDRequest) MustDo() *GetAddressbookPartiesIDContactsContactIDAccountsContactAccountIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID products product ID events API
*/
func (a *GetAddressbookPartiesIDProductsProductIDEventsRequest) Do() (*GetAddressbookPartiesIDProductsProductIDEventsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDProductsProductIDEvents",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/products/{product_id}/events",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDProductsProductIDEventsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDProductsProductIDEventsOK), nil

}

func (a *GetAddressbookPartiesIDProductsProductIDEventsRequest) MustDo() *GetAddressbookPartiesIDProductsProductIDEventsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get addressbook parties ID products product ID events event ID API
*/
func (a *GetAddressbookPartiesIDProductsProductIDEventsEventIDRequest) Do() (*GetAddressbookPartiesIDProductsProductIDEventsEventIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAddressbookPartiesIDProductsProductIDEventsEventID",
		Method:             "GET",
		PathPattern:        "/addressbook/parties/{id}/products/{product_id}/events/{event_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAddressbookPartiesIDProductsProductIDEventsEventIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddressbookPartiesIDProductsProductIDEventsEventIDOK), nil

}

func (a *GetAddressbookPartiesIDProductsProductIDEventsEventIDRequest) MustDo() *GetAddressbookPartiesIDProductsProductIDEventsEventIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
patch addressbook parties ID actors party actor ID API
*/
func (a *PatchAddressbookPartiesIDActorsPartyActorIDRequest) Do() (*PatchAddressbookPartiesIDActorsPartyActorIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAddressbookPartiesIDActorsPartyActorID",
		Method:             "PATCH",
		PathPattern:        "/addressbook/parties/{id}/actors/{party_actor_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PatchAddressbookPartiesIDActorsPartyActorIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchAddressbookPartiesIDActorsPartyActorIDOK), nil

}

func (a *PatchAddressbookPartiesIDActorsPartyActorIDRequest) MustDo() *PatchAddressbookPartiesIDActorsPartyActorIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post addressbook parties API
*/
func (a *PostAddressbookPartiesRequest) Do() (*PostAddressbookPartiesCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAddressbookParties",
		Method:             "POST",
		PathPattern:        "/addressbook/parties",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostAddressbookPartiesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAddressbookPartiesCreated), nil

}

func (a *PostAddressbookPartiesRequest) MustDo() *PostAddressbookPartiesCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post addressbook parties ID actors API
*/
func (a *PostAddressbookPartiesIDActorsRequest) Do() (*PostAddressbookPartiesIDActorsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAddressbookPartiesIDActors",
		Method:             "POST",
		PathPattern:        "/addressbook/parties/{id}/actors",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostAddressbookPartiesIDActorsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAddressbookPartiesIDActorsCreated), nil

}

func (a *PostAddressbookPartiesIDActorsRequest) MustDo() *PostAddressbookPartiesIDActorsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post addressbook parties ID contacts API
*/
func (a *PostAddressbookPartiesIDContactsRequest) Do() (*PostAddressbookPartiesIDContactsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAddressbookPartiesIDContacts",
		Method:             "POST",
		PathPattern:        "/addressbook/parties/{id}/contacts",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostAddressbookPartiesIDContactsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAddressbookPartiesIDContactsCreated), nil

}

func (a *PostAddressbookPartiesIDContactsRequest) MustDo() *PostAddressbookPartiesIDContactsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post addressbook parties ID contacts contact ID accounts API
*/
func (a *PostAddressbookPartiesIDContactsContactIDAccountsRequest) Do() (*PostAddressbookPartiesIDContactsContactIDAccountsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAddressbookPartiesIDContactsContactIDAccounts",
		Method:             "POST",
		PathPattern:        "/addressbook/parties/{id}/contacts/{contact_id}/accounts",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostAddressbookPartiesIDContactsContactIDAccountsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAddressbookPartiesIDContactsContactIDAccountsCreated), nil

}

func (a *PostAddressbookPartiesIDContactsContactIDAccountsRequest) MustDo() *PostAddressbookPartiesIDContactsContactIDAccountsCreated {
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
