package piholeapi

import (
	"fmt"
	"net/url"
)

const (
	dnsEndpoint = "/config/dns/hosts"
)

type DNSConfigResponse struct {
	Hosts            []string `json:"hosts"`
	Upstreams        []string `json:"upstreams"`
	CNAMEdeepInspect bool     `json:"CNAMEdeepInspect"`
}

type PiholeConfigResponse struct {
	DNS DNSConfigResponse `json:"dns"`
}

type ConfigResponse struct {
	Config PiholeConfigResponse `json:"config"`
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
