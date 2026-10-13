package piholeapi

import "context"

const (
	actionEndpoint     = "/action"
	gravityEndpoint    = actionEndpoint + "/gravity"
	restartDNSEndpoint = actionEndpoint + "/restartdns"
)

// UpdateGravity updates Pi-hole's gravity database.
func (c *Client) UpdateGravity(ctx context.Context) error {
	return c.post(ctx, gravityEndpoint, nil, nil)
}

// RestartDNS restarts Pi-hole's DNS resolver.
func (c *Client) RestartDNS(ctx context.Context) error {
	return c.post(ctx, restartDNSEndpoint, nil, nil)
}
