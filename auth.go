package piholeapi

import (
	"errors"
	"fmt"
)

const (
	AuthEndpoint = "/auth"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginSession struct {
	Valid    bool   `json:"valid"`
	Totp     bool   `json:"totp"`
	Sid      string `json:"sid"`
	Csrf     string `json:"csrf"`
	Validity int    `json:"validity"`
	Message  string `json:"message"`
}

type LoginResponse struct {
	Session LoginSession `json:"session"`
	Took    float64      `json:"took"`
}

func (c *Client) Auth(password string) error {
	var LoginResp LoginResponse

	err := c.PostRequest(
		AuthEndpoint,
		LoginRequest{
			Password: password,
		},
		&LoginResp,
	)
	if err != nil {
		return err
	}

	session := LoginResp.Session
	if session.Totp {
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

	if session.Sid == "" {
		return errors.New("authentication failed: missing session ID")
	}

	c.Sid = session.Sid

	return nil
}
