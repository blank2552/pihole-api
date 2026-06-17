package piholeapi

import (
	"fmt"
	"net/url"
	"slices"
)

const (
	ConfigEndpoint = "/config"
)

type DNSConfig struct {
	Hosts     []string `json:"hosts,omitempty"`
	Upstreams []string `json:"upstreams,omitempty"`
}

type Config struct {
	DNS DNSConfig `json:"dns"`
}

type PiHoleConfigResponse struct {
	Config Config `json:"config"`
}

func (c *Client) DNSHostExists(ip, hostname string) (bool, error) {
	resp, err := c.GetPiholeConfig()
	if err != nil {
		return false, err
	}

	want := ip + " " + hostname

	return slices.Contains(resp.Config.DNS.Hosts, want), nil
}

func (c *Client) AddDNSEntry(address, hostname string) error {
	endpoint := fmt.Sprintf(
		ConfigEndpoint+"/dns/hosts/%s",
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

func (c *Client) AddUpstream(address string) error {
	endpoint := ConfigEndpoint + "/dns/upstreams/" + address

	if err := c.put(
		endpoint,
		nil,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetPiholeConfig() (*PiHoleConfigResponse, error) {
	var resp PiHoleConfigResponse

	if err := c.get(
		ConfigEndpoint,
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
