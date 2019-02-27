package federation

import (
	"net/http"
	"net/url"

	"github.com/test/go/clients/testhorizon"
	"github.com/test/go/clients/testtoml"
	proto "github.com/test/go/protocols/federation"
)

// FederationResponseMaxSize is the maximum size of response from a federation server
const FederationResponseMaxSize = 100 * 1024

// DefaultTestNetClient is a default federation client for testnet
var DefaultTestNetClient = &Client{
	HTTP:        http.DefaultClient,
	TestHorizon:     testhorizon.DefaultTestNetClient,
	TestTOML: testtoml.DefaultClient,
}

// DefaultPublicNetClient is a default federation client for pubnet
var DefaultPublicNetClient = &Client{
	HTTP:        http.DefaultClient,
	TestHorizon:     testhorizon.DefaultPublicNetClient,
	TestTOML: testtoml.DefaultClient,
}

// Client represents a client that is capable of resolving a federation request
// using the internet.
type Client struct {
	TestTOML TestTOML
	HTTP        HTTP
	TestHorizon     TestHorizon
	AllowHTTP   bool
}

type ClientInterface interface {
	LookupByAddress(addy string) (*proto.NameResponse, error)
	LookupByAccountID(aid string) (*proto.IDResponse, error)
	ForwardRequest(domain string, fields url.Values) (*proto.NameResponse, error)
}

// TestHorizon represents a testhorizon client that can be consulted for data when
// needed as part of the federation protocol
type TestHorizon interface {
	HomeDomainForAccount(aid string) (string, error)
}

// HTTP represents the http client that a federation client uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// TestTOML represents a client that can resolve a given domain name to
// test.toml file.  The response is used to find the federation server that a
// query should be made against.
type TestTOML interface {
	GetTestToml(domain string) (*testtoml.Response, error)
}

// confirm interface conformity
var _ TestTOML = testtoml.DefaultClient
var _ HTTP = http.DefaultClient
var _ ClientInterface = &Client{}
