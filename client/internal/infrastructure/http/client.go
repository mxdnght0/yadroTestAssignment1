package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"yadroTestAssignment/client/internal/contracts"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) contracts.DNSClient {
	return &Client{baseURL: baseURL}
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type apiErrorResponse struct {
	Error apiError `json:"error"`
}

func parseAPIError(body []byte) (string, error) {
	var e apiErrorResponse
	if err := json.Unmarshal(body, &e); err != nil {
		return "", err
	}
	return e.Error.Code, nil
}

func (c *Client) doRequest(method, path string) ([]byte, int, error) {
	req, err := http.NewRequest(method, c.baseURL+path, nil)
	if err != nil {
		return nil, 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot reach server at %s: %w", c.baseURL, err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

func (c *Client) Add(ip string) error {
	path := "/dns?" + url.Values{"dns": {ip}}.Encode()
	body, status, err := c.doRequest(http.MethodPost, path)
	if err != nil {
		return err
	}
	if status == http.StatusCreated {
		return nil
	}

	code, _ := parseAPIError(body)
	switch code {
	case "DNS_INVALID":
		return ErrInvalidIP
	case "DNS_ALREADY_EXISTS":
		return ErrAlreadyExists
	default:
		return fmt.Errorf("unexpected server error (HTTP %d)", status)
	}
}

func (c *Client) Delete(ip string) error {
	path := "/dns?" + url.Values{"dns": {ip}}.Encode()
	body, status, err := c.doRequest(http.MethodDelete, path)
	if err != nil {
		return err
	}
	if status == http.StatusOK {
		return nil
	}

	code, _ := parseAPIError(body)
	switch code {
	case "DNS_INVALID":
		return ErrInvalidIP
	case "DNS_NOT_FOUND":
		return ErrNotFound
	case "FILE_NOT_FOUND":
		return ErrFileNotCreated
	default:
		return fmt.Errorf("unexpected server error (HTTP %d)", status)
	}
}

func (c *Client) GetAll() ([]string, error) {
	body, status, err := c.doRequest(http.MethodGet, "/dns")
	if err != nil {
		return nil, err
	}
	if status == http.StatusNotFound {
		return []string{}, nil
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("unexpected server error (HTTP %d)", status)
	}

	var result struct {
		Lines []string `json:"dns_lines"`
	}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Lines, nil
}
