package piholeapi

func (c *Client) History() (any, error) {
	var Resp map[string]any

	err := c.GetRequest(
		"/history",
		&Resp,
	)
	if err != nil {
		return err
	}

	return resp, nil
}
