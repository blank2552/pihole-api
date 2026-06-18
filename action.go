package piholeapi

const (
	actionEndpoint     = "/action"
	gravityEndpoint    = actionEndpoint + "/gravity"
	restartDNSEndpoint = actionEndpoint + "/restartdns"
)

func (c *Client) UpdateGravity() error {
	return c.post(gravityEndpoint, nil, nil)
}

func (c *Client) RestartDNS() error {
	return c.post(restartDNSEndpoint, nil, nil)
}
