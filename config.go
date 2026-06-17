package piholeapi

import (
	"fmt"
	"net/url"
)

const (
	dnsEndpoint = "/config/dns/hosts"
)

type HostsConfigResponse struct {
	Hosts []string `json:"hosts"`
}

type DNSConfigResponse struct {
	DNS       HostsConfigResponse `json:"dns"`
	Upstreams []string            `json:"upstreams"`
}

type ConfigResponse struct {
	Config DNSConfigResponse `json:"config"`
}

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

func (c *Client) GetDNSEntries() (*ConfigResponse, error) {
	var resp ConfigResponse

	if err := c.get(
		dnsEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
