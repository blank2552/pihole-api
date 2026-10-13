# pihole-api

A Go client for the Pi-hole v6 REST API.

## Status

The client is ready for early use, but its public API may change before v1.0.0.

## Requirements

- Go 1.21 or newer
- Pi-hole v6

## Installation

Install the latest release:

```sh
go get github.com/blank2552/pihole-api@v0.1.0
```

## Quick start

The base URL must include Pi-hole's `/api` path. Set `verifyTLS` to `true` for
normal certificate verification.

```go
package main

import (
	"context"
	"fmt"
	"log"

	piholeapi "github.com/blank2552/pihole-api"
)

func main() {
	ctx := context.Background()
	client := piholeapi.NewClient("https://pi.hole/api", true)

	if err := client.Login(ctx, "your-password"); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Logout(ctx); err != nil {
			log.Printf("logout: %v", err)
		}
	}()

	history, err := client.GetHistory(ctx)
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

if err := client.Login(ctx, password); err != nil {
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
| Statistics | `GetHistory`, `GetClientHistory`, `GetDatabaseHistory`, `GetQueries` |

### Create a domain-list entry

```go
err := client.CreateListEntry(ctx, piholeapi.CreateListEntryRequest{
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
if err := client.UpdateGravity(ctx); err != nil {
	return err
}
```

## Limitations

- TOTP authentication is not supported.
- Compatibility may vary between Pi-hole FTL versions.

## Security

Set `verifyTLS` to `true` in normal use. Replacing `Client.HTTPClient` also
replaces its default redirect protection.

## License

This project is available under the [MIT License](LICENSE).
