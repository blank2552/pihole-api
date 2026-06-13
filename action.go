package piholeapi

const (
	gravityEndpoint    = "/action/gravity"
	restartDNSEndpoint = "/action/restartdns"
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

func (c *Client) RestartDNS() error {
	err := c.post(
		restartDNSEndpoint,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
