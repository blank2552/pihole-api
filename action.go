package piholeapi

const (
	GravityEndpoint		= "/action/gravity"
)

func (c *Client) UpdateGravity() error {
	err := c.PostRequest(
		GravityEndpoint,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
