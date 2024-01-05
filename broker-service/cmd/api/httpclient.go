package main

import (
	"net/http"
)

// HTTPClient is a reusable HTTP client with custom configuration.
type HTTPClient struct {
	Client *http.Client
}

// NewHTTPClient creates a new instance of HTTPClient with custom configurations.
func NewHTTPClient() *HTTPClient {
	client := &http.Client{
		// Timeout: timeout,
		// You can customize other settings like Transport, CheckRedirect, etc.
	}

	return &HTTPClient{Client: client}
}
