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


type ClientsEntry struct {
	Name string `json:"name"`
	Total int `json:"total"`
}

type ClientsHistoryEntry struct {
	Timestamp int `json:"timestamp"`
	Data map[string]int `json:"data"`
}

type HistoryClientsResponse struct {
	Clients map[string]ClientsEntry `json:"clients"`
	History []ClientsHistoryEntry `json:"history"`
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

	return &Resp, nil
}

func (c *Client) HistoryClients(n int) (*HistoryClientsResponse, error) {
	var Resp HistoryClientsResponse

	err := c.GetRequest(
		"/history/clients",
		&Resp,
	)
	if err != nil {
		return nil, err
	}

	return &Resp, nil
}
