package kongfig

import "github.com/hbagdi/go-kong/kong"

// Config represents a declarative configuration generated using kongfig of Kong < 0.14
type Config struct {
	Apis         []API         `json:"apis"`
	Consumers    []Consumer    `json:"consumers"`
	Plugins      []Plugin      `json:"plugins"`
	Upstreams    []interface{} `json:"upstreams"`
	Certificates []interface{} `json:"certificates"`
}

// API represents an API object in Kong < 0.14
type API struct {
	Name       string        `json:"name"`
	Plugins    []Plugin      `json:"plugins"`
	Attributes APIAttributes `json:"attributes"`
}

// APIAttributes represents the attributes of a API
type APIAttributes struct {
	Hosts                  []string `json:"hosts"`
	Uris                   []string `json:"uris"`
	StripURI               bool     `json:"strip_uri"`
	PreserveHost           bool     `json:"preserve_host"`
	UpstreamURL            string   `json:"upstream_url"`
	Retries                int      `json:"retries"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout"`
	HTTPSOnly              bool     `json:"https_only"`
	HTTPIfTerminated       bool     `json:"http_if_terminated"`
	Methods                []string `json:"methods"`
}

// Plugin represents a Plugin in Kong < 0.14
type Plugin struct {
	Name       string           `json:"name"`
	Attributes PluginAttributes `json:"attributes"`
}

// PluginAttributes represents the attributes of a Kong Plugin < 0.14
type PluginAttributes struct {
	Enabled bool               `json:"enabled"`
	Config  kong.Configuration `json:"config"`
}

// Consumer represents a consumer in Kong < 0.14
type Consumer struct {
	ID       *string `json:"id,omitempty" yaml:"id,omitempty"`
	Username string  `json:"username"`
	CustomID *string `json:"custom_id,omitempty"`
	// Acls        []string     `json:"acls"`
	Credentials []Credential `json:"credentials"`
}

// Credential represents a consumer credential
type Credential struct {
	Name       string               `json:"name"`
	Attributes CredentialAttributes `json:"attributes"`
}

// CredentialAttributes represents the Credential attributes
type CredentialAttributes struct {
	RSAPublicKey *string `json:"rsa_public_key,omitempty"`
	Algorithm    string  `json:"algorithm,omitempty"`
	Key          string  `json:"key"`
	Secret       *string `json:"secret,omitempty"`
}
