package http

import (
	"fmt"
	"net/http"
)

/*
The golang stdlib http.Client uses the http.DefaultTransport by defualt (and
subsequently http.Get(), etc).

That DefaultTransport uses the stdlib crypto/tls.
To override usage of the stdlib crypto/tls, just set this OpenSSLRoundTripper
as the DefaultTransport for your application. Likely in an init() somewhere.

  import "net/http"
  func init() {
    http.DefaultTransport = &OpenSSLRoundTripper{}
  }

*/
type OpenSSLRoundTripper struct {
  // a default RoundTripper for all non-https schemes
  // ( this defaults to http.DefaultTransport )
	DefaultTransport http.RoundTripper
}

var (
	errorNilRequest  = fmt.Errorf("http: nil Request")
	errorNilURL      = fmt.Errorf("http: nil Request.URL")
	errorNilHeader   = fmt.Errorf("http: nil Request.Header")
	defaultTransport = http.DefaultTransport
)

func (osrt *OpenSSLRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, errorNilRequest
	}
	if req.URL == nil {
		req.Body.Close()
		return nil, errorNilURL
	}

	// short circuit to most use-cases
	if req.URL.Scheme != "https" {
		if osrt.DefaultTransport != nil {
			return osrt.DefaultTransport.RoundTrip(req)
		}
		return defaultTransport.RoundTrip(req)
	}

	if req.Header == nil {
		req.Body.Close()
		return nil, errorNilHeader
	}

	// XXX

	return nil, nil
}
