package piholeapi

const (
	historyEndpoint = "/history"
)

// HistoryEntry contains aggregate DNS query counts for a point in time.
type HistoryEntry struct {
	Timestamp int64 `json:"timestamp"`
	Total     int   `json:"total"`
	Cached    int   `json:"cached"`
	Blocked   int   `json:"blocked"`
	Forwarded int   `json:"forwarded"`
}

// HistoryResponse contains Pi-hole's DNS query history.
type HistoryResponse struct {
	History []HistoryEntry `json:"history"`
	Took    float64        `json:"took"`
}

// GetHistory returns Pi-hole's aggregate DNS query history.
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
