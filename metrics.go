package piholeapi

import (
	"net/url"
	"strconv"
	"time"
)

const (
	historyEndpoint         = "/history"
	clientHistoryEndpoint   = historyEndpoint + "/clients"
	databaseHistoryEndpoint = historyEndpoint + "/database"
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

// ClientHistoryClient contains aggregate query information for a client.
type ClientHistoryClient struct {
	Name  *string `json:"name"`
	Total int     `json:"total"`
}

// ClientHistoryEntry contains per-client query counts for a point in time.
type ClientHistoryEntry struct {
	Timestamp int64          `json:"timestamp"`
	Data      map[string]int `json:"data"`
}

// DatabaseHistoryEntry
type DatabaseHistoryEntry struct {
	Timestamp int64 `json:"timestamp"`
	Total     int   `json:"total"`
	Cached    int   `json:"cached"`
	Blocked   int   `json:"blocked"`
	Forwarded int   `json:"forwarded"`
}

// DatabaseHistoryResponse contains Pi-hole's Database query history.
type DatabaseHistoryResponse struct {
	History []DatabaseHistoryEntry `json:"history"`
	Took    float64                `json:"took"`
}

type DatabaseHistoryOptions struct {
	// From limits results from this time.
	From time.Time `json:"time"`

	// Until limits results until this time.
	Until time.Time `json:"until"`
}

// ClientHistoryResponse contains DNS query history grouped by client. Client
// map keys are identifiers assigned by Pi-hole.
type ClientHistoryResponse struct {
	Clients map[string]ClientHistoryClient `json:"clients"`
	History []ClientHistoryEntry           `json:"history"`
	Took    float64                        `json:"took"`
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

// GetClientHistory returns Pi-hole's DNS query history grouped by client.
func (c *Client) GetClientHistory() (*ClientHistoryResponse, error) {
	var resp ClientHistoryResponse

	if err := c.get(
		clientHistoryEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetDatabaseHistory(
	options DatabaseHistoryOptions,
) (*DatabaseHistoryResponse, error) {
	values := url.Values{}
	values.Set("from", strconv.FormatInt(options.From.Unix(), 10))
	values.Set("until", strconv.FormatInt(options.Until.Unix(), 10))

	endpoint := databaseHistoryEndpoint + "?" + values.Encode()

	var resp DatabaseHistoryResponse
	if err := c.get(
		endpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
