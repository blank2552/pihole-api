package piholeapi

import "net/url"

const (
	listsEndpoint = "/lists"
)

// ListEntry describes an entry in a Pi-hole domain list.
type ListEntry struct {
	Address      string  `json:"address"`
	Comment      *string `json:"comment,omitempty"`
	Groups       []int   `json:"groups"`
	Enabled      bool    `json:"enabled"`
	ID           int     `json:"id"`
	DateAdded    int64   `json:"date_added"`
	DateModified int64   `json:"date_modified"`
}

// ListsResponse contains Pi-hole domain-list entries.
type ListsResponse struct {
	Lists []ListEntry `json:"lists"`
	Took  float64     `json:"took"`
}

// CreateListEntryRequest contains the values used to create a Pi-hole domain-list entry.
type CreateListEntryRequest struct {
	Address string `json:"address"`
	Comment string `json:"comment,omitempty"`
	Groups  []int  `json:"groups"`
	Enabled bool   `json:"enabled"`
}

// GetListEntries returns all Pi-hole domain-list entries.
func (c *Client) GetListEntries() (*ListsResponse, error) {
	var resp ListsResponse

	if err := c.get(
		listsEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetListEntriesByAddress returns Pi-hole domain-list entries matching address.
func (c *Client) GetListEntriesByAddress(address string) (*ListsResponse, error) {
	var resp ListsResponse

	if err := c.get(
		listsEndpoint+"/"+url.PathEscape(address),
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateListEntry creates a Pi-hole domain-list entry.
func (c *Client) CreateListEntry(l CreateListEntryRequest) error {
	if err := c.post(
		listsEndpoint+"/"+url.PathEscape(l.Address),
		&l,
		nil,
	); err != nil {
		return err
	}

	return nil
}
