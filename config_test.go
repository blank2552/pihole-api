package piholeapi

import (
	"net/http"
	"testing"
)

func TestDeleteUpstream(t *testing.T) {
	client := testClient(t, http.MethodDelete, "/api/config/dns/upstreams/1.1.1.1%2353")

	if err := client.DeleteUpstream("1.1.1.1#53"); err != nil {
		t.Fatalf("DeleteUpstream() error = %v", err)
	}
}
