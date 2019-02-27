package testtoml

import (
	"fmt"
	"io"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/test/go/address"
	"github.com/test/go/support/errors"
)

// GetTestToml returns test.toml file for a given domain
func (c *Client) GetTestToml(domain string) (resp *Response, err error) {
	var hresp *http.Response
	hresp, err = c.HTTP.Get(c.url(domain))
	if err != nil {
		err = errors.Wrap(err, "http request errored")
		return
	}
	defer hresp.Body.Close()

	if !(hresp.StatusCode >= 200 && hresp.StatusCode < 300) {
		err = errors.New("http request failed with non-200 status code")
		return
	}

	limitReader := io.LimitReader(hresp.Body, TestTomlMaxSize)
	_, err = toml.DecodeReader(limitReader, &resp)

	// There is one corner case not handled here: response is exactly
	// TestTomlMaxSize long and is incorrect toml. Check discussion:
	// https://github.com/test/go/pull/24#discussion_r89909696
	if err != nil && limitReader.(*io.LimitedReader).N == 0 {
		err = errors.Errorf("test.toml response exceeds %d bytes limit", TestTomlMaxSize)
		return
	}

	if err != nil {
		err = errors.Wrap(err, "toml decode failed")
		return
	}

	return
}

// GetTestTomlByAddress returns test.toml file of a domain fetched from a
// given address
func (c *Client) GetTestTomlByAddress(addy string) (*Response, error) {
	_, domain, err := address.Split(addy)
	if err != nil {
		return nil, errors.Wrap(err, "parse address failed")
	}

	return c.GetTestToml(domain)
}

// url returns the appropriate url to load for resolving domain's test.toml
// file
func (c *Client) url(domain string) string {
	var scheme string

	if c.UseHTTP {
		scheme = "http"
	} else {
		scheme = "https"
	}

	return fmt.Sprintf("%s://%s%s", scheme, domain, WellKnownPath)
}
