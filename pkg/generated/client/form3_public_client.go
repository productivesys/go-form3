// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/a_c_e"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/accounts"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/addressbook"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/audit"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/balances"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/claims"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/direct_debits"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/fundingrequest_api"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/fx_api"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/mandates"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/oauth2"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/organisations"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/participants"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/payments"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/query_api"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/reports"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/roles"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/scheme_messages"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/subscriptions"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/users"
)

// Default form3 public HTTP client.
// var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.form3.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

/*
// NewHTTPClient creates a new form3 public HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Form3Public {
  return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new form3 public HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Form3Public {
  // ensure nullable parameters have default
  if cfg == nil {
    cfg = DefaultTransportConfig()
  }

  // create transport and client
  transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
  return New(transport, formats)
}
*/

// New creates a new form3 public client
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Form3Public {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Form3Public)
	cli.Transport = transport

	cli.ACE = a_c_e.New(transport, formats, defaults)

	cli.Accounts = accounts.New(transport, formats, defaults)

	cli.Addressbook = addressbook.New(transport, formats, defaults)

	cli.Audit = audit.New(transport, formats, defaults)

	cli.Balances = balances.New(transport, formats, defaults)

	cli.Claims = claims.New(transport, formats, defaults)

	cli.DirectDebits = direct_debits.New(transport, formats, defaults)

	cli.FundingrequestAPI = fundingrequest_api.New(transport, formats, defaults)

	cli.FxAPI = fx_api.New(transport, formats, defaults)

	cli.Mandates = mandates.New(transport, formats, defaults)

	cli.Oauth2 = oauth2.New(transport, formats, defaults)

	cli.Organisations = organisations.New(transport, formats, defaults)

	cli.Participants = participants.New(transport, formats, defaults)

	cli.Payments = payments.New(transport, formats, defaults)

	cli.QueryAPI = query_api.New(transport, formats, defaults)

	cli.Reports = reports.New(transport, formats, defaults)

	cli.Roles = roles.New(transport, formats, defaults)

	cli.SchemeMessages = scheme_messages.New(transport, formats, defaults)

	cli.Subscriptions = subscriptions.New(transport, formats, defaults)

	cli.Users = users.New(transport, formats, defaults)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Form3Public is a client for form3 public
type Form3Public struct {
	ACE *a_c_e.Client

	Accounts *accounts.Client

	Addressbook *addressbook.Client

	Audit *audit.Client

	Balances *balances.Client

	Claims *claims.Client

	DirectDebits *direct_debits.Client

	FundingrequestAPI *fundingrequest_api.Client

	FxAPI *fx_api.Client

	Mandates *mandates.Client

	Oauth2 *oauth2.Client

	Organisations *organisations.Client

	Participants *participants.Client

	Payments *payments.Client

	QueryAPI *query_api.Client

	Reports *reports.Client

	Roles *roles.Client

	SchemeMessages *scheme_messages.Client

	Subscriptions *subscriptions.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Form3Public) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.ACE.SetTransport(transport)

	c.Accounts.SetTransport(transport)

	c.Addressbook.SetTransport(transport)

	c.Audit.SetTransport(transport)

	c.Balances.SetTransport(transport)

	c.Claims.SetTransport(transport)

	c.DirectDebits.SetTransport(transport)

	c.FundingrequestAPI.SetTransport(transport)

	c.FxAPI.SetTransport(transport)

	c.Mandates.SetTransport(transport)

	c.Oauth2.SetTransport(transport)

	c.Organisations.SetTransport(transport)

	c.Participants.SetTransport(transport)

	c.Payments.SetTransport(transport)

	c.QueryAPI.SetTransport(transport)

	c.Reports.SetTransport(transport)

	c.Roles.SetTransport(transport)

	c.SchemeMessages.SetTransport(transport)

	c.Subscriptions.SetTransport(transport)

	c.Users.SetTransport(transport)

}
