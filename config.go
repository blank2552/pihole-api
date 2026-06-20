package piholeapi

import (
	"fmt"
	"net/url"
	"slices"
)

const (
	configEndpoint       = "/config"
	dnsHostsEndpoint     = "/dns/hosts"
	dnsUpstreamsEndpoint = "/dns/upstreams"
)

type DNSConfig struct {
	Hosts     []string `json:"hosts,omitempty"`
	Upstreams []string `json:"upstreams,omitempty"`
}

type Config struct {
	DNS DNSConfig `json:"dns"`
}

type ConfigResponse struct {
	Config Config  `json:"config"`
	Took   float64 `json:"took"`
}

func dnsHostEntry(address, hostname string) string {
	return address + " " + hostname
}

func (c *Client) DNSHostExists(address, hostname string) (bool, error) {
	resp, err := c.GetPiholeConfig()
	if err != nil {
		return false, err
	}

	want := dnsHostEntry(address, hostname)

	return slices.Contains(resp.Config.DNS.Hosts, want), nil
}

func (c *Client) AddDNSEntry(address, hostname string) error {
	endpoint := fmt.Sprintf(
		configEndpoint+dnsHostsEndpoint+"/%s",
		url.PathEscape(dnsHostEntry(address, hostname)),
	)

	return c.put(endpoint, nil, nil)
}

func (c *Client) DeleteDNSEntry(address, hostname string) error {
	endpoint := fmt.Sprintf(
		configEndpoint+dnsHostsEndpoint+"/%s",
		url.PathEscape(dnsHostEntry(address, hostname)),
	)

	return c.delete(endpoint, nil, nil)
}

func (c *Client) AddUpstream(address string) error {
	endpoint := fmt.Sprintf(
		configEndpoint+dnsUpstreamsEndpoint+"/%s",
		url.PathEscape(address),
	)

	return c.put(endpoint, nil, nil)
}

func (c *Client) GetPiholeConfig() (*ConfigResponse, error) {
	var resp ConfigResponse

	if err := c.get(
		configEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
