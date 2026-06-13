# pihole-api

Go client for the Pi-hole API.

This package currently supports authentication, query history, adding local DNS host entries, and triggering a gravity update.

## Installation

```sh
go get github.com/blank2552/pihole-api
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	piholeapi "github.com/blank2552/pihole-api"
)

func main() {
	client := piholeapi.NewClient("https://pi.hole/api", true)

	if err := client.Login("your-password"); err != nil {
		log.Fatal(err)
	}

	history, err := client.GetHistory()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("history entries: %d\n", len(history.History))
}
```

## Client

Create a client with `NewClient`:

```go
client := piholeapi.NewClient("https://pi.hole/api", true)
```

The second argument controls TLS certificate verification:

- `true` verifies TLS certificates.
- `false` disables TLS verification. Use this only for trusted local development or self-signed local Pi-hole installations.

After a successful login, the client stores the session ID and sends it on later requests using the `X-FTL-SID` header.

## Authentication

```go
err := client.Login("your-password")
```

`Login` returns an error if authentication fails, if the API requires TOTP, or if the API response does not include a session ID.

## History

```go
history, err := client.GetHistory()
if err != nil {
	return err
}

for _, entry := range history.History {
	fmt.Println(entry.Timestamp, entry.Total, entry.Blocked)
}
```

## DNS Entries

Add a local DNS host entry:

```go
err := client.AddDNSEntry("192.168.1.10", "server.local")
```

## Gravity

Trigger a gravity update:

```go
err := client.UpdateGravity()
```

## Supported Methods

- `NewClient(baseURL string, verifyTLS bool) *Client`
- `Login(password string) error`
- `GetHistory() (*HistoryResponse, error)`
- `AddDNSEntry(address, hostname string) error`
- `UpdateGravity() error`

## Notes

- The base URL should include the Pi-hole API path, for example `https://pi.hole/api`.
- API responses with HTTP status code `400` or higher are returned as errors.
- TOTP authentication is detected, but entering a TOTP code is not currently implemented.
