package piholeapi

import (
	"errors"
	"fmt"
)

const (
	authEndpoint = "/auth"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginSession struct {
	Valid    bool   `json:"valid"`
	TOTP     bool   `json:"totp"`
	SID      string `json:"sid"`
	CSRF     string `json:"csrf"`
	Validity int    `json:"validity"`
	Message  string `json:"message"`
}

type LoginResponse struct {
	Session LoginSession `json:"session"`
	Took    float64      `json:"took"`
}

func (c *Client) Login(password string) error {
	var loginResp LoginResponse

	err := c.post(
		authEndpoint,
		LoginRequest{
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
			return fmt.Errorf("authentication requires TOTP: %s", session.Message)
		}

		return errors.New("authentication requires TOTP")
	}

	if !session.Valid {
		if session.Message != "" {
			return fmt.Errorf("authentication failed: %s", session.Message)
		}

		return errors.New("authentication failed")
	}

	if session.SID == "" {
		return errors.New("authentication failed: missing session ID")
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

	return nil
}
