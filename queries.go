package piholeapi

const (
	queryEndpoint = "/queries"
)

type client struct {
	IP   string `json:"ip"`
	Name string `json:"name"`
}

type reply struct {
	Type string `json:"type"`
	Time int64  `json:"time"`
}

type ede struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type queryEntry struct {
	ID        int    `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Domain    string `json:"domain"`
	CNAME     string `json:"cname"`
	Status    string `json:"status"`
	Client    client `json:"client"`
	DNSSEC    string `json:"dnssec"`
	Reply     reply  `json:"reply"`
	ListID    int    `json:"list_id"`
	Upstream  string `json:"upstream"`
	Ede       ede    `json:"ede"`
}

type queriesResponse struct {
	Queries []queryEntry `json:"queries"`
	Took    float64      `json:"took"`
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
