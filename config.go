package piholeapi

import (
	"fmt"
	"net/url"
	"slices"
)

const (
	dnsEndpoint = "/config/dns/hosts"
)

type DNSConfig struct {
	Hosts []string `json:"hosts,omitempty"`
}

type Config struct {
	DNS DNSConfig `json:"dns"`
}

type PiHoleConfigResponse struct {
	Config Config `json:"config"`
}

func (c *Client) DNSHostExists(ip, hostname string) (bool, error) {
	cfg, err := c.GetDNSEntries()
	if err != nil {
		return false, err
	}

	want := ip + " " + hostname

	return slices.Contains(cfg.Config.DNS.Hosts, want), nil
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

func (c *Client) GetDNSEntries() (*PiHoleConfigResponse, error) {
	var resp PiHoleConfigResponse

	if err := c.get(
		dnsEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
