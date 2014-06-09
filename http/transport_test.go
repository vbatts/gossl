package http

import (
	"testing"
)

func TestRoundTripper(t *testing.T) {
	o := OpenSSLRoundTripper{}
	r, err := o.RoundTrip(nil)
	if err != errorNilRequest {
		t.Errorf("this should have raised an error")
	}
	_ = r
}
