// File: factory.go

/*
	Factory is a factory for the core client.
*/

package core

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

// Factory is a factory for the core client.
type Factory struct {
	Builder *Builder
}

// NewFactory returns a new factory with a new builder.
func NewFactory() *Factory {
	return &Factory{
		Builder: &Builder{},
	}
}

// NewFactoryWithBuilder returns a new factory with a provided core builder.
func NewFactoryWithBuilder(pBuilder *Builder) *Factory {
	return &Factory{
		Builder: pBuilder,
	}
}

// CreateClient creates a new core client with a new logger and a new http client.
func (f *Factory) CreateClient() *Client {
	var pLogger *zap.Logger

	pLogger, _ = zap.NewProduction()

	f.Builder.WithLogger(pLogger)
	f.Builder.WithHttpClient(&http.Client{})

	return f.Builder.GetInstance()
}

// CreateWithLogger creates a new core client with a logger.
func (f *Factory) CreateWithLogger(logger LoggerInterface) *Client {
	f.Builder.WithLogger(logger)
	f.Builder.WithHttpClient(&http.Client{})

	return f.Builder.GetInstance()
}

// CreateWithHttpClient creates a new core client with a http client.
func (f *Factory) CreateWithHttpClient(pClient *http.Client) *Client {
	var err error
	var pLogger *zap.Logger

	if pLogger, err = zap.NewProduction(); err != nil {
		fmt.Println(err.Error())
	}

	f.Builder.WithLogger(pLogger)
	f.Builder.WithHttpClient(pClient)

	return f.Builder.GetInstance()
}

// CreateWithLoggerAndHttpClient creates a new core client with a logger and a http client.
func (f *Factory) CreateWithLoggerAndHttpClient(loggerInterface LoggerInterface, pClient *http.Client) *Client {
	f.Builder.WithLogger(loggerInterface)
	f.Builder.WithHttpClient(pClient)

	return f.Builder.GetInstance()
}
