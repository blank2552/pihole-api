# pihole-api

Go client for the Pi-hole API.


## Installation

```sh
go get codeberg.org/blank2552/go-pihole-api-client
```

## Example usage

### Authenticate

```go
client := piholeapi.NewClient("https://pi.hole/api", true)
if err := client.Login("your-password"); err != nil {
	return err
}
defer client.Logout()
```

### Get lists

```go
resp, err := client.GetLists()
if err != nil {
    panic(err)
}

fmt.Printf("%+v", resp)
```

### Update gravity

```go
if err := client.UpdateGravity(); err != nil {
    panic(err)
}
```

The base URL should include the Pi-hole API path, for example `https://pi.hole/api`.
