package piholeapi

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetClientHistory(t *testing.T) {
	client := NewClient("http://pihole.test/api", true)
	client.HTTPClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want %q", r.Method, http.MethodGet)
		}
		if got, want := r.URL.RequestURI(), "/api/history/clients"; got != want {
			t.Errorf("request URI = %q, want %q", got, want)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader(`{
				"clients":{"192.0.2.1":{"name":"laptop","total":12}},
				"history":[{"timestamp":1700000000,"data":{"192.0.2.1":3}}],
				"took":0.001
			}`)),
		}, nil
	})

	resp, err := client.GetClientHistory()
	if err != nil {
		t.Fatalf("GetClientHistory() error = %v", err)
	}

	entry := resp.Clients["192.0.2.1"]
	if entry.Name == nil || *entry.Name != "laptop" || entry.Total != 12 {
		t.Errorf("client entry = %#v", entry)
	}
	if got := resp.History[0].Data["192.0.2.1"]; got != 3 {
		t.Errorf("history count = %d, want 3", got)
	}
}
