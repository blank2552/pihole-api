package piholeapi

const (
	listsEndpoint = "/lists"
)

type ListEntry struct {
	Address      string `json:"address"`
	Comment      string `json:"comment"`
	Groups       []int  `json:"groups"`
	Enabled      bool   `json:"enabled"`
	ID           int    `json:"id"`
	DateAdded    int    `json:"date_added"`
	DateModified int    `json:"date_modified"`
}

type ListsResponse struct {
	Lists []ListEntry `json:"lists"`
}

func (c *Client) GetLists() (*ListsResponse, error) {
	var resp ListsResponse

	err := c.get(
		listsEndpoint,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetListsByAddress(address string) (*ListsResponse, error) {
	var resp ListsResponse

	err := c.get(
		listsEndpoint+"/"+address,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
