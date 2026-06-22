package piholeapi

import "net/url"

const (
	listsEndpoint = "/lists"
)

type ListEntry struct {
	Address      string  `json:"address"`
	Comment      *string `json:"comment,omitempty"`
	Groups       []int   `json:"groups"`
	Enabled      bool    `json:"enabled"`
	ID           int     `json:"id"`
	DateAdded    int     `json:"date_added"`
	DateModified int     `json:"date_modified"`
}

type ListsResponse struct {
	Lists []ListEntry `json:"lists"`
	Took  int         `json:"took"`
}

type ListCreate struct {
	Address string `json:"address"`
	Comment string `json:"comment,omitempty"`
	Groups  []int  `json:"groups"`
	Enabled bool   `json:"enabled"`
}

func (c *Client) GetLists() (*ListsResponse, error) {
	var resp ListsResponse

	if err := c.get(
		listsEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetListsByAddress(address string) (*ListsResponse, error) {
	var resp ListsResponse

	if err := c.get(
		listsEndpoint+"/"+url.PathEscape(address),
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CreateList(l ListCreate) error {
	if err := c.post(
		listsEndpoint+"/"+url.PathEscape(l.Address),
		&l,
		nil,
	); err != nil {
		return err
	}

	return nil
}
