package core

import (
	"net/http"
)

type Builder struct {
	Logger     LoggerInterface
	HttpClient *http.Client
}

func (b *Builder) WithLogger(loggerInterface LoggerInterface) *Builder {
	b.Logger = loggerInterface
	return b
}

func (b *Builder) WithHttpClient(pClient *http.Client) *Builder {
	b.HttpClient = pClient
	return b
}

func (b *Builder) GetInstance() *Client {
	pClient := NewClient()

	pClient.Logger = b.Logger
	pClient.HttpClient = b.HttpClient

	return pClient
}
