package piholeapi

const (
	gravityEndpoint = "/action/gravity"
)

func (c *Client) UpdateGravity() error {
	err := c.post(
		gravityEndpoint,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
