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

// DNSConfig contains Pi-hole's custom DNS hosts and upstream resolvers.
type DNSConfig struct {
	Hosts     []string `json:"hosts,omitempty"`
	Upstreams []string `json:"upstreams,omitempty"`
}

// Config contains the supported Pi-hole configuration sections.
type Config struct {
	DNS DNSConfig `json:"dns"`
}

// ConfigResponse contains Pi-hole's configuration response.
type ConfigResponse struct {
	Config Config  `json:"config"`
	Took   float64 `json:"took"`
}

func dnsHostEntry(address, hostname string) string {
	return address + " " + hostname
}

// DNSHostExists reports whether an exact address and hostname entry exists.
func (c *Client) DNSHostExists(address, hostname string) (bool, error) {
	resp, err := c.GetConfig()
	if err != nil {
		return false, err
	}

	want := dnsHostEntry(address, hostname)

	return slices.Contains(resp.Config.DNS.Hosts, want), nil
}

// AddDNSEntry adds an address and hostname entry to Pi-hole's custom DNS hosts.
func (c *Client) AddDNSEntry(address, hostname string) error {
	endpoint := fmt.Sprintf(
		configEndpoint+dnsHostsEndpoint+"/%s",
		url.PathEscape(dnsHostEntry(address, hostname)),
	)

	return c.put(endpoint, nil, nil)
}

// DeleteDNSEntry removes an address and hostname entry from Pi-hole's custom DNS hosts.
func (c *Client) DeleteDNSEntry(address, hostname string) error {
	endpoint := fmt.Sprintf(
		configEndpoint+dnsHostsEndpoint+"/%s",
		url.PathEscape(dnsHostEntry(address, hostname)),
	)

	return c.delete(endpoint, nil, nil)
}

// AddUpstream adds a DNS upstream resolver to Pi-hole's configuration.
func (c *Client) AddUpstream(address string) error {
	endpoint := fmt.Sprintf(
		configEndpoint+dnsUpstreamsEndpoint+"/%s",
		url.PathEscape(address),
	)

	return c.put(endpoint, nil, nil)
}

// DeleteUpstream removes a DNS upstream resolver from Pi-hole's configuration.
func (c *Client) DeleteUpstream(address string) error {
	endpoint := fmt.Sprintf(
		configEndpoint+dnsUpstreamsEndpoint+"/%s",
		url.PathEscape(address),
	)

	return c.delete(endpoint, nil, nil)
}

// GetConfig returns Pi-hole's current configuration.
func (c *Client) GetConfig() (*ConfigResponse, error) {
	var resp ConfigResponse

	if err := c.get(
		configEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
