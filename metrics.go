package piholeapi

func (c *Client) Hist() (any, error) {
	var Resp map[string]any

	err := c.GetRequest(
		"/history",
		&Resp,
	)
	if err != nil {
		return nil, err
	}

	return Resp, nil
}
