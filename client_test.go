package piholeapi

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func testClient(t *testing.T, wantMethod, wantURI string) *Client {
	t.Helper()

	client := NewClient("http://pihole.test/api", true)
	client.HTTPClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.Method != wantMethod {
			t.Errorf("method = %q, want %q", r.Method, wantMethod)
		}
		if got := r.URL.RequestURI(); got != wantURI {
			t.Errorf("request URI = %q, want %q", got, wantURI)
		}

		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader("")),
			Header:     make(http.Header),
		}, nil
	})

	return client
}
