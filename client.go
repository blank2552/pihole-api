package pihole-api

import (
	"encoding/json"
	"net/http"
	"bytes"
)

type Client struct {
	URL string
	Password string `json="password"`
	Client *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
		Client: &http.Client{},
	}
}

func (c *Client) GetRequest(path string, result any) error {
	req, err := http.NewRequest(
		http.MethodGet,
		c.URL+path,
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("api returned %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *Client) PostRequest(path string, body any, result any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.URL+path,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type": "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer Body.Close()
	
	if resp.StatusCode >= 400 {
		return fmt.Errorf("api returned %s", resp.Status)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}
