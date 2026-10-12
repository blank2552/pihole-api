package piholeapi

import (
	"net/url"
	"strconv"
	"time"
)

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

// QueryOptions contains optional filters and pagination values for querying
// Pi-hole's DNS query log. Zero-valued fields use Pi-hole's defaults.
type QueryOptions struct {
	// From limits results to queries at or after this time.
	From time.Time
	// Until limits results to queries at or before this time.
	Until time.Time
	// Domain filters results by the queried domain.
	Domain string
	// ClientIP filters results by the requesting client's IP address.
	ClientIP string
	// ClientName filters results by the requesting client's hostname.
	ClientName string
	// Status filters results by Pi-hole query status.
	Status string
	// Length sets the maximum number of queries returned.
	Length int
	// Start sets the pagination offset.
	Start int
}

// GetQueries returns DNS queries recorded by Pi-hole using opts as filters and
// pagination values. Zero-valued options use Pi-hole's defaults.
func (c *Client) GetQueries(opts QueryOptions) (*QueriesResponse, error) {
	values := url.Values{}

	if !opts.From.IsZero() {
		values.Set("from", strconv.FormatInt(opts.From.Unix(), 10))
	}

	if !opts.Until.IsZero() {
		values.Set("until", strconv.FormatInt(opts.Until.Unix(), 10))
	}

	if opts.Domain != "" {
		values.Set("domain", opts.Domain)
	}

	if opts.ClientIP != "" {
		values.Set("client_ip", opts.ClientIP)
	}

	if opts.ClientName != "" {
		values.Set("client_name", opts.ClientName)
	}

	if opts.Status != "" {
		values.Set("status", opts.Status)
	}

	if opts.Length > 0 {
		values.Set("length", strconv.Itoa(opts.Length))
	}

	if opts.Start > 0 {
		values.Set("start", strconv.Itoa(opts.Start))
	}

	endpoint := queryEndpoint
	if encoded := values.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}

	var resp QueriesResponse
	if err := c.get(
		endpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
