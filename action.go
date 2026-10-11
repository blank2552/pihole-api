package piholeapi

const (
	actionEndpoint     = "/action"
	gravityEndpoint    = actionEndpoint + "/gravity"
	restartDNSEndpoint = actionEndpoint + "/restartdns"
)

// UpdateGravity updates Pi-hole's gravity database.
func (c *Client) UpdateGravity() error {
	return c.post(gravityEndpoint, nil, nil)
}

// RestartDNS restarts Pi-hole's DNS resolver.
func (c *Client) RestartDNS() error {
	return c.post(restartDNSEndpoint, nil, nil)
}
