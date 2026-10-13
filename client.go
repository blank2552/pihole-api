package piholeapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ErrRedirectNotAllowed indicates that the Pi-hole API attempted to redirect a request.
var ErrRedirectNotAllowed = errors.New("pihole API redirect not allowed")

// Client is a client for the Pi-hole API.
type Client struct {
	baseURL string

	authMu sync.Mutex
	sidMu  sync.RWMutex
	sid    string

	// HTTPClient performs API requests and may be replaced before first use.
	HTTPClient *http.Client
}

// NewClient creates a Pi-hole API client for baseURL.
// TLS certificate verification is disabled when verifyTLS is false.
func NewClient(baseURL string, verifyTLS bool) *Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()

	if !verifyTLS {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec -- explicitly requested by caller
		}
	}

	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		HTTPClient: &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
			CheckRedirect: func(
				req *http.Request,
				via []*http.Request,
			) error {
				return ErrRedirectNotAllowed
			},
		},
	}
}

func (c *Client) sessionID() string {
	c.sidMu.RLock()
	defer c.sidMu.RUnlock()

	return c.sid
}

func (c *Client) setSessionID(sid string) {
	c.sidMu.Lock()
	defer c.sidMu.Unlock()

	c.sid = sid
}

func (c *Client) get(ctx context.Context, path string, result any) error {
	return c.doRequest(ctx, http.MethodGet, path, nil, result)
}

func (c *Client) post(ctx context.Context, path string, body any, result any) error {
	return c.doRequest(ctx, http.MethodPost, path, body, result)
}

func (c *Client) put(ctx context.Context, path string, body any, result any) error {
	return c.doRequest(ctx, http.MethodPut, path, body, result)
}

func (c *Client) delete(ctx context.Context, path string, body any, result any) error {
	return c.doRequest(ctx, http.MethodDelete, path, body, result)
}

func (c *Client) doRequest(ctx context.Context, method, path string, body any, result any) error {
	requestBody, err := encodeRequestBody(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		requestBody,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	if sid := c.sessionID(); sid != "" {
		req.Header.Set("X-FTL-SID", sid)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	const maxErrorBodyBytes = 64 * 1024

	if resp.StatusCode >= http.StatusBadRequest {
		data, err := io.ReadAll(
			io.LimitReader(resp.Body, maxErrorBodyBytes),
		)
		if err != nil {
			return fmt.Errorf(
				"read pihole api error response for %s %s: %w",
				method,
				path,
				err,
			)
		}

		return &APIError{
			Method:     method,
			Path:       path,
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    strings.TrimSpace(string(data)),
		}
	}

	if result == nil {
		return nil
	}

	contentType := resp.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return fmt.Errorf(
			"pihole API returned invalid content-type %q for %s %s: %w",
			contentType,
			method,
			path,
			err,
		)
	}

	if mediaType != "application/json" {
		return fmt.Errorf(
			"pihole API returned unexpected content-type %q for %s %s",
			contentType,
			method,
			path,
		)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf(
			"decode pihole response for %s %s: %w",
			method,
			path,
			err,
		)
	}

	return nil
}

func encodeRequestBody(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(data), nil
}
