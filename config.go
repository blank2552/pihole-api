package piholeapi

import (
	"fmt"
	"net/url"
)

const (
	DNSEndpoint		= "/config/dns/hosts"
)


func (c *Client) AddDNSEntry(hostname, address string) error {
        
		endpoint := fmt.Sprintf(
			DNSEndpoint+"%s",
			url.PathEscape(address+" "+hostname),
		) 

		fmt.Println(c.URL+endpoint)

        err := c.PutRequest(
                DNSEndpoint+endpoint,
                nil,
				nil,
        )
        if err != nil {
                return err
        }

        return nil
}
