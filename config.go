package piholeapi

import (
	"fmt"
	"net/url"
)

const (
	DNSEndpoint		= "/config/dns.hosts"
)


func (c *Client) AddDNSEntry(address, hostname string) error {
        
		endpoint := fmt.Sprintf(
			DNSEndpoint+"/%s",
			url.PathEscape(address+" "+hostname),
		) 

		fmt.Println(endpoint)

        err := c.PutRequest(
                endpoint,
                nil,
				nil,
        )
        if err != nil {
                return err
        }

        return nil
}
