# pihole-api

Go client for the Pi-hole API.

This package currently supports authentication, session listing, query history, local DNS configuration, list management, triggering a gravity update, and restarting DNS.

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
- `Logout() error`
- `ListSessions() (*AuthSessionsResponse, error)`
- `GetHistory() (*HistoryResponse, error)`
- `GetPiholeConfig() (*ConfigResponse, error)`
- `DNSHostExists(address, hostname string) (bool, error)`
- `AddDNSEntry(address, hostname string) error`
- `AddUpstream(address string) error`
- `GetLists() (*ListsResponse, error)`
- `GetListsByAddress(address string) (*ListsResponse, error)`
- `CreateList(l ListCreate) error`
- `UpdateGravity() error`
- `RestartDNS() error`

## Notes

- API responses with HTTP status code `400` or higher are returned as errors.
- TOTP authentication is detected, but entering a TOTP code is not currently implemented.
