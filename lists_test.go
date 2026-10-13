package piholeapi

import (
	"context"
	"net/http"
	"testing"
)

func TestCreateListEntry(t *testing.T) {
	client := testClient(t, http.MethodPost, "/api/lists?type=block")

	err := client.CreateListEntry(context.Background(), CreateListEntryRequest{
		Address: "https://example.com/list.txt",
		Type:    "block",
		Comment: "example",
		Groups:  []int{0},
		Enabled: true,
	})
	if err != nil {
		t.Fatalf("CreateListEntry() error = %v", err)
	}
}

func TestUpdateListEntry(t *testing.T) {
	client := testClient(t, http.MethodPut, "/api/lists/https:%2F%2Fexample.com%2Flist.txt?type=block")

	err := client.UpdateListEntry(context.Background(), "https://example.com/list.txt", UpdateListEntryRequest{
		Type:    "block",
		Comment: "example",
		Groups:  []int{0},
		Enabled: true,
	})
	if err != nil {
		t.Fatalf("UpdateListEntry() error = %v", err)
	}
}

func TestDeleteListEntry(t *testing.T) {
	client := testClient(t, http.MethodDelete, "/api/lists/https:%2F%2Fexample.com%2Flist.txt?type=block")

	if err := client.DeleteListEntry(context.Background(), "https://example.com/list.txt", "block"); err != nil {
		t.Fatalf("DeleteListEntry() error = %v", err)
	}
}
