package piholeapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Client struct {
	baseURL    string
	sidMu      sync.RWMutex
	sid        string
	HTTPClient *http.Client
}

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

func (c *Client) get(path string, result any) error {
	return c.doRequest(http.MethodGet, path, nil, result)
}

func (c *Client) post(path string, body any, result any) error {
	return c.doRequest(http.MethodPost, path, body, result)
}

func (c *Client) put(path string, body any, result any) error {
	return c.doRequest(http.MethodPut, path, body, result)
}

func (c *Client) delete(path string, body any, result any) error {
	return c.doRequest(http.MethodDelete, path, body, result)
}

func (c *Client) doRequest(method, path string, body any, result any) error {
	requestBody, err := encodeRequestBody(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
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

	if resp.StatusCode >= 400 {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("api returned %s", resp.Status)
		}

		message := strings.TrimSpace(string(data))
		if message != "" {
			return fmt.Errorf("api returned %s: %s", resp.Status, message)
		}

		return fmt.Errorf("api returned %s", resp.Status)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
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
