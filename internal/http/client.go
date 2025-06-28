package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

type Response struct {
	StatusCode   int
	Headers      map[string]string
	Body         string
	ResponseTime time.Duration
	Error        error
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) SendRequest(req Request) Response {
	start := time.Now()

	// Prepare request body
	var reqBody io.Reader
	if req.Body != "" && (req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH") {
		reqBody = bytes.NewBufferString(req.Body)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(context.Background(), req.Method, req.URL, reqBody)
	if err != nil {
		return Response{Error: fmt.Errorf("failed to create request: %w", err)}
	}

	// Set default headers
	httpReq.Header.Set("User-Agent", "Quest/1.0")
	if reqBody != nil && httpReq.Header.Get("Content-Type") == "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	// Set custom headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Send request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return Response{Error: fmt.Errorf("request failed: %w", err)}
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{Error: fmt.Errorf("failed to read response body: %w", err)}
	}

	// Parse response headers
	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = strings.Join(values, ", ")
	}

	return Response{
		StatusCode:   resp.StatusCode,
		Headers:      headers,
		Body:         string(body),
		ResponseTime: time.Since(start),
	}
}

func FormatResponse(body string) string {
	// Try to format as JSON
	var jsonData interface{}
	if err := json.Unmarshal([]byte(body), &jsonData); err == nil {
		formatted, err := json.MarshalIndent(jsonData, "", "  ")
		if err == nil {
			return string(formatted)
		}
	}

	return body
}

func ValidateURL(url string) error {
	if url == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return fmt.Errorf("URL must start with http:// or https://")
	}

	return nil
}
