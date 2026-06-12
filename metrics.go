package piholeapi

type HistoryEntry struct {
	Timestamp int `json:"timestamp"`
	Total	int `json:"total"`
	Cached	int `json:"cached"`
	Blocked	int `json:"blocked"`
	Forwarded	int `json:"forwarded"`
}

type HistoryResponse struct {
	History []HistoryEntry `json:"history"`
	Took float64 `json:"took"`
}


func (c *Client) History() (*HistoryResponse, error) {
	var Resp HistoryResponse

	err := c.GetRequest(
		"/history",
		&Resp,
	)
	if err != nil {
		return nil, err
	}

	return Resp, nil
}
