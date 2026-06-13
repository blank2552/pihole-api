package piholeapi

import (
	"fmt"
	"net/url"
)

const (
	dnsEndpoint = "/config/dns/hosts"
)

func (c *Client) AddDNSEntry(address, hostname string) error {

	endpoint := fmt.Sprintf(
		dnsEndpoint+"/%s",
		url.PathEscape(address+" "+hostname),
	)

	err := c.put(
		endpoint,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
