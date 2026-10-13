package piholeapi

import (
	"context"
	"net/http"
	"testing"
)

func TestDeleteUpstream(t *testing.T) {
	client := testClient(t, http.MethodDelete, "/api/config/dns/upstreams/1.1.1.1%2353")

	if err := client.DeleteUpstream(context.Background(), "1.1.1.1#53"); err != nil {
		t.Fatalf("DeleteUpstream() error = %v", err)
	}
}
