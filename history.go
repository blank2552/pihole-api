package piholeapi

const (
	historyEndpoint = "/history"
)

type HistoryEntry struct {
	Timestamp int64 `json:"timestamp"`
	Total     int   `json:"total"`
	Cached    int   `json:"cached"`
	Blocked   int   `json:"blocked"`
	Forwarded int   `json:"forwarded"`
}

type HistoryResponse struct {
	History []HistoryEntry `json:"history"`
	Took    float64        `json:"took"`
}

func (c *Client) GetHistory() (*HistoryResponse, error) {
	var resp HistoryResponse

	if err := c.get(
		historyEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
