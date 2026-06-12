package piholeapi

import (
	"fmt"
	"net/url"
)

const (
	DNSEndpoint		= "/config/dns"
)


func (c *Client) AddDNSEntry(hostname string, address string) error {
        
		endpoint := fmt.Sprintf(
			DNSEndpoint+"/hosts/%s",
			url.PathEscape(address+" "+hostname),
		) 

        err := c.GetRequest(
                endpoint,
                nil,
        )
        if err != nil {
                return err
        }

        return nil
}
