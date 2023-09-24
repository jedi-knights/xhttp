package core

import (
	"net/http"
)

// Builder is a builder for the core client.
type Builder struct {
	Logger     LoggerInterface
	HttpClient *http.Client
}

// WithLogger sets the logger.
func (b *Builder) WithLogger(loggerInterface LoggerInterface) *Builder {
	b.Logger = loggerInterface
	return b
}

// WithHttpClient sets the http client.
func (b *Builder) WithHttpClient(pClient *http.Client) *Builder {
	b.HttpClient = pClient
	return b
}

// GetInstance returns a new core client with a logger and a http client.
func (b *Builder) GetInstance() *Client {
	pClient := NewClient()

	pClient.Logger = b.Logger
	pClient.HttpClient = b.HttpClient

	return pClient
}
