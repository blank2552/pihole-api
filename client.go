package piholeapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL    string
	SID        string
	HTTPClient *http.Client
}

func NewClient(baseURL string, verifyTLS bool) *Client {
	tr := &http.Transport{}

	if !verifyTLS {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Transport: tr,
		},
	}
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

func (c *Client) doRequest(method, path string, body any, result any) error {
	requestBody, err := encodeRequestBody(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		method,
		c.BaseURL+path,
		requestBody,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	if c.SID != "" {
		req.Header.Set("X-FTL-SID", c.SID)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("api returned %s", resp.Status)
	}

	if result != nil {
		fmt.Println(json.NewDecoder(resp.Body).Decode(result))
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
