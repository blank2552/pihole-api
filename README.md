# pihole-api

Go client for the Pi-hole API.

This package currently supports authentication, query history, adding local DNS host entries, and triggering a gravity update.

## Installation

```sh
go get github.com/blank2552/pihole-api
```

## Usage

```go
client := piholeapi.NewClient("https://pi.hole/api", true)
err := client.Login("your-password")
```

The base URL should include the Pi-hole API path, for example `https://pi.hole/api`.

The second `NewClient` argument controls TLS verification. Pass `false` only for trusted local development or self-signed local Pi-hole installations.

## Examples

- `examples/history`: authenticate and print query history.
- `examples/add_dns_entry`: add a local DNS host entry.
- `examples/update_gravity`: trigger a gravity update.

## Supported Methods

- `NewClient(baseURL string, verifyTLS bool) *Client`
- `Login(password string) error`
- `GetHistory() (*HistoryResponse, error)`
- `AddDNSEntry(address, hostname string) error`
- `UpdateGravity() error`

## Notes

- API responses with HTTP status code `400` or higher are returned as errors.
- TOTP authentication is detected, but entering a TOTP code is not currently implemented.
