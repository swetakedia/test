package testtoml

import (
	"net/http"
	"strings"
	"testing"

	"github.com/test/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://test.org/.well-known/test.toml", c.url("test.org"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://test.org/.well-known/test.toml", c.url("test.org"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://test.org/.well-known/test.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetTestToml("test.org")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// test.toml exceeds limit
	h.
		On("GET", "https://toobig.org/.well-known/test.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation`+strings.Repeat("0", TestTomlMaxSize)+`"`,
		)
	stoml, err = c.GetTestToml("toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "test.toml response exceeds")
	}

	// not found
	h.
		On("GET", "https://missing.org/.well-known/test.toml").
		ReturnNotFound()
	stoml, err = c.GetTestToml("missing.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://json.org/.well-known/test.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetTestToml("json.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
