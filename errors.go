package piholeapi

import "fmt"

// APIError represents an unsuccessful HTTP response from the Pi-hole API.
type APIError struct {
	Method     string // HTTP method used for the request.
	Path       string // API path requested.
	Status     string // HTTP response status, such as "401 Unauthorized".
	StatusCode int    // HTTP response status code.
	Message    string // Bounded response body returned by the API.
}

// Error returns a human-readable description of the unsuccessful API response.
func (e APIError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf(
			"pihole API %s %s returned %s",
			e.Method,
			e.Path,
			e.Status,
		)
	}

	return fmt.Sprintf(
		"pihole API %s %s returned %s: %s",
		e.Method,
		e.Path,
		e.Status,
		e.Message,
	)
}
