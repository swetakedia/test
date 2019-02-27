package testtoml

import "net/http"

// TestTomlMaxSize is the maximum size of test.toml file
const TestTomlMaxSize = 5 * 1024

// WellKnownPath represents the url path at which the test.toml file should
// exist to conform to the federation protocol.
const WellKnownPath = "/.well-known/test.toml"

// DefaultClient is a default client using the default parameters
var DefaultClient = &Client{HTTP: http.DefaultClient}

// Client represents a client that is capable of resolving a Test.toml file
// using the internet.
type Client struct {
	// HTTP is the http client used when resolving a Test.toml file
	HTTP HTTP

	// UseHTTP forces the client to resolve against servers using plain HTTP.
	// Useful for debugging.
	UseHTTP bool
}

type ClientInterface interface {
	GetTestToml(domain string) (*Response, error)
	GetTestTomlByAddress(addy string) (*Response, error)
}

// HTTP represents the http client that a stellertoml resolver uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// Response represents the results of successfully resolving a test.toml file
type Response struct {
	AuthServer       string `toml:"AUTH_SERVER"`
	FederationServer string `toml:"FEDERATION_SERVER"`
	EncryptionKey    string `toml:"ENCRYPTION_KEY"`
	SigningKey       string `toml:"SIGNING_KEY"`
}

// GetTestToml returns test.toml file for a given domain
func GetTestToml(domain string) (*Response, error) {
	return DefaultClient.GetTestToml(domain)
}

// GetTestTomlByAddress returns test.toml file of a domain fetched from a
// given address
func GetTestTomlByAddress(addy string) (*Response, error) {
	return DefaultClient.GetTestTomlByAddress(addy)
}
