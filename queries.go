package piholeapi

const (
	queryEndpoint = "/queries"
)

// QueryClient identifies the client that issued a DNS query.
type QueryClient struct {
	IP   string  `json:"ip"`
	Name *string `json:"name,omitempty"`
}

// QueryReply describes Pi-hole's reply to a DNS query.
type QueryReply struct {
	Type *string `json:"type,omitempty"`
	Time float64 `json:"time"`
}

// QueryEDE contains extended DNS error information for a query.
type QueryEDE struct {
	Code int     `json:"code"`
	Text *string `json:"text,omitempty"`
}

// QueryEntry describes a DNS query recorded by Pi-hole.
type QueryEntry struct {
	ID        int         `json:"id"`
	Timestamp *int64      `json:"timestamp"`
	Type      string      `json:"type"`
	Domain    string      `json:"domain"`
	CNAME     *string     `json:"cname,omitempty"`
	Status    *string     `json:"status,omitempty"`
	Client    QueryClient `json:"client"`
	DNSSEC    string      `json:"dnssec"`
	Reply     QueryReply  `json:"reply"`
	ListID    *int        `json:"list_id,omitempty"`
	Upstream  *string     `json:"upstream,omitempty"`
	Ede       QueryEDE    `json:"ede"`
}

// QueriesResponse contains DNS query records and pagination metadata.
type QueriesResponse struct {
	Queries               []QueryEntry `json:"queries"`
	Cursor                int          `json:"cursor"`
	RecordsTotal          int          `json:"recordsTotal"`
	RecordsFiltered       int          `json:"recordsFiltered"`
	Draw                  int          `json:"draw"`
	EarliestTimestamp     int64        `json:"earliest_timestamp"`
	EarliestTimestampDisk int64        `json:"earliest_timestamp_disk"`
	Took                  float64      `json:"took"`
}

// GetQueries returns DNS queries recorded by Pi-hole.
func (c *Client) GetQueries() (*QueriesResponse, error) {
	var resp QueriesResponse

	if err := c.get(
		queryEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
