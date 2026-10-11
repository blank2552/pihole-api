package piholeapi

import (
	"errors"
	"fmt"
)

const (
	authEndpoint = "/auth"
)

var (
	// ErrTOTPRequired indicates that authentication requires a TOTP code.
	ErrTOTPRequired = errors.New("authentication requires TOTP")
	// ErrAuthenticationFailed indicates that Pi-hole rejected the credentials.
	ErrAuthenticationFailed = errors.New("authentication failed")
	// ErrMissingSessionID indicates that authentication succeeded without a session ID.
	ErrMissingSessionID = errors.New("authentication failed: missing session ID")
)

type authRequest struct {
	Password string `json:"password"`
}

type authSession struct {
	Valid    bool   `json:"valid"`
	TOTP     bool   `json:"totp"`
	SID      string `json:"sid"`
	CSRF     string `json:"csrf"`
	Validity int    `json:"validity"`
	Message  string `json:"message"`
}

type authResponse struct {
	Session authSession `json:"session"`
	Took    float64     `json:"took"`
}

// AuthSessionTLS describes the TLS properties of an authenticated session.
type AuthSessionTLS struct {
	Login bool `json:"login"`
	Mixed bool `json:"mixed"`
}

// AuthSessionInfo describes an active Pi-hole API session.
type AuthSessionInfo struct {
	ID             int            `json:"id"`
	CurrentSession bool           `json:"current_session"`
	Valid          bool           `json:"valid"`
	TLS            AuthSessionTLS `json:"tls"`
	App            bool           `json:"app"`
	CLI            bool           `json:"cli"`
	LoginAt        int64          `json:"login_at"`
	LastActive     int64          `json:"last_active"`
	ValidUntil     int64          `json:"valid_until"`
	RemoteAddr     string         `json:"remote_addr"`
	UserAgent      *string        `json:"user_agent,omitempty"`
	XForwardedFor  *string        `json:"x_forwarded_for,omitempty"`
}

// AuthSessionsResponse contains the active Pi-hole API sessions.
type AuthSessionsResponse struct {
	Sessions []AuthSessionInfo `json:"sessions"`
}

// Login authenticates with Pi-hole and stores the returned session ID.
func (c *Client) Login(password string) error {
	c.setSessionID("")

	var loginResp authResponse

	if err := c.post(
		authEndpoint,
		authRequest{
			Password: password,
		},
		&loginResp,
	); err != nil {
		return err
	}

	session := loginResp.Session
	if session.TOTP {
		if session.Message != "" {
			return fmt.Errorf("%w: %s", ErrTOTPRequired, session.Message)
		}

		return ErrTOTPRequired
	}

	if !session.Valid {
		if session.Message != "" {
			return fmt.Errorf("%w: %s", ErrAuthenticationFailed, session.Message)
		}

		return ErrAuthenticationFailed
	}

	if session.SID == "" {
		return ErrMissingSessionID
	}

	c.setSessionID(session.SID)

	return nil
}

// Logout invalidates the current Pi-hole session and clears its session ID.
func (c *Client) Logout() error {
	if err := c.delete(
		authEndpoint,
		nil,
		nil,
	); err != nil {
		return err
	}

	c.setSessionID("")

	return nil
}

// ListSessions returns the active Pi-hole API sessions.
func (c *Client) ListSessions() (*AuthSessionsResponse, error) {
	var resp AuthSessionsResponse

	if err := c.get(
		authEndpoint+"/sessions",
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
