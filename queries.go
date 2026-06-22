package piholeapi

const (
	queryEndpoint = "/queries"
)

type client struct {
	IP   string `json:"ip"`
	Name string `json:"name,omitempty"`
}

type reply struct {
	Type string  `json:"type,omitempty"`
	Time float64 `json:"time"`
}

type ede struct {
	Code int    `json:"code"`
	Text string `json:"text,omitempty"`
}

type queryEntry struct {
	ID        int    `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Domain    string `json:"domain"`
	CNAME     string `json:"cname,omitempty"`
	Status    string `json:"status,omitempty"`
	Client    client `json:"client"`
	DNSSEC    string `json:"dnssec"`
	Reply     reply  `json:"reply"`
	ListID    int    `json:"list_id,omitempty"`
	Upstream  string `json:"upstream,omitempty"`
	Ede       ede    `json:"ede"`
}

type queriesResponse struct {
	Queries               []queryEntry `json:"queries"`
	Cursor                int          `json:"cursor"`
	RecordsTotal          int          `json:"recordsTotal"`
	RecordsFiltered       int          `json:"recordsFiltered"`
	Draw                  int          `json:"draw"`
	EarliestTimestamp     int64        `json:"earliest_timestamp"`
	EarliestTimestampDisk int64        `json:"earliest_timestamp_disk"`
	Took                  float64      `json:"took"`
}

func (c *Client) GetQueries() (*queriesResponse, error) {
	var resp queriesResponse

	if err := c.get(
		queryEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
