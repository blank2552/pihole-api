package piholeapi

import (
	"errors"
	"fmt"
)

const (
	authEndpoint = "/auth"
)

var (
	ErrTOTPRequired         = errors.New("authentication requires TOTP")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrMissingSessionID     = errors.New("authentication failed: missing session ID")
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

type authCurrentSession struct {
	ID             int  `json:"id"`
	CurrentSession bool `json:"current_session"`
	Valid          bool `json:"valid"`
	TLS            struct {
		Login bool `json:"login"`
		Mixed bool `json:"mixed"`
	} `json:"tls"`
	APP        bool   `json:"app"`
	CLI        bool   `json:"cli"`
	LoginAt    int    `json:"login_at"`
	LastActive int    `json:"last_active"`
	ValidUntil int    `json:"valid_until"`
	RemoteAddr string `json:"remote_addr"`
	UserAgent  string `json:"user_agent,omitempty"`
	XFF        string `json:"x_forwarded_for,omitempty"`
}

type authListSessions struct {
	Sessions []authCurrentSession `json:"sessions"`
}

func (c *Client) Login(password string) error {
	c.SID = ""

	var loginResp authResponse

	err := c.post(
		authEndpoint,
		authRequest{
			Password: password,
		},
		&loginResp,
	)
	if err != nil {
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

	c.SID = session.SID

	return nil
}

func (c *Client) Logout() error {
	if err := c.delete(
		authEndpoint,
		nil,
		nil,
	); err != nil {
		return err
	}

	c.SID = ""

	return nil
}

func (c *Client) ListSessions() (*authListSessions, error) {
	var resp authListSessions

	if err := c.get(
		authEndpoint+"/sessions",
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
