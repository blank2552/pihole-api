package piholeapi

const (
	ListsEndpoint = "/lists"
)

type ListObject struct {
	Address      string `json:"address"`
	Comment      string `json:"comment"`
	Groups       []int  `json:"groups"`
	Enabled      bool   `json:"enabled"`
	ID           int    `json:"id"`
	DateAdded    int    `json:"date_added"`
	DateModified int    `json:"date_modified"`
}

type ListsResponse struct {
	Lists []ListObject `json:"lists"`
}

func (c *Client) Lists() (*ListsResponse, error) {
	var resp ListsResponse

	err := c.get(
		ListsEndpoint,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
