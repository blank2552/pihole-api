# pihole-api

Go client for the Pi-hole API.


## Installation

```sh
go get github.com/blank2552/pihole-api
```

## Usage

```go
client := piholeapi.NewClient("https://pi.hole/api", true)
if err := client.Login("your-password"); err != nil {
	return err
}
defer client.Logout()
```

The base URL should include the Pi-hole API path, for example `https://pi.hole/api`.
