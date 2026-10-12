# pihole-api

A Go client for the Pi-hole v6 REST API.

## Requirements

- Go 1.21 or newer
- Pi-hole v6

## Installation

The current development version has not been released as a new tag yet. Install
it directly from `main`:

```sh
go get github.com/blank2552/pihole-api@main
```

## Quick start

The base URL must include Pi-hole's `/api` path. Set `verifyTLS` to `true` for
normal certificate verification.

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

	defer func() {
		if err := client.Logout(); err != nil {
			log.Printf("logout: %v", err)
		}
	}()

	history, err := client.GetHistory()
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range history.History {
		fmt.Printf("%+v\n", entry)
	}
}
```

## Error handling

HTTP error responses are returned as `*APIError`. Use `errors.As` to inspect
the status without parsing an error string:

```go
var apiErr *piholeapi.APIError

if err := client.Login(password); err != nil {
	if errors.As(err, &apiErr) {
		log.Printf(
			"Pi-hole returned HTTP %d: %s",
			apiErr.StatusCode,
			apiErr.Message,
		)
	}

	return err
}
```

## API overview

| Area | Methods |
| --- | --- |
| Authentication | `Login`, `Logout`, `ListSessions` |
| Actions | `UpdateGravity`, `RestartDNS` |
| Configuration | `GetConfig`, `DNSHostExists`, `AddDNSEntry`, `DeleteDNSEntry`, `AddUpstream`, `DeleteUpstream` |
| Domain lists | `GetListEntries`, `GetListEntriesByAddress`, `CreateListEntry`, `UpdateListEntry`, `DeleteListEntry` |
| Statistics | `GetHistory`, `GetClientHistory`, `GetQueries` |

### Create a domain-list entry

```go
	err := client.CreateListEntry(piholeapi.CreateListEntryRequest{
	Address: "example.com",
	Type:    "block",
	Comment: "Example blocklist entry",
	Groups:  []int{0},
	Enabled: true,
})
if err != nil {
	return err
}
```

### Update gravity

```go
if err := client.UpdateGravity(); err != nil {
	return err
}
```

## Limitations

- TOTP authentication is detected but cannot yet be completed by the client.
- Request-specific contexts are not yet supported.

## License

This project is available under the [MIT License](LICENSE).
