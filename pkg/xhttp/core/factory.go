package core

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type Factory struct {
	Builder *Builder
}

func NewFactory() *Factory {
	return &Factory{
		Builder: &Builder{},
	}
}

func NewFactoryWithBuilder(pBuilder *Builder) *Factory {
	return &Factory{
		Builder: pBuilder,
	}
}

func (f *Factory) CreateClient() *Client {
	var pLogger *zap.Logger

	pLogger, _ = zap.NewProduction()

	f.Builder.WithLogger(pLogger)
	f.Builder.WithHttpClient(&http.Client{})

	return f.Builder.GetInstance()
}

func (f *Factory) CreateWithLogger(logger LoggerInterface) *Client {
	f.Builder.WithLogger(logger)
	f.Builder.WithHttpClient(&http.Client{})

	return f.Builder.GetInstance()
}

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

func (f *Factory) CreateWithLoggerAndHttpClient(loggerInterface LoggerInterface, pClient *http.Client) *Client {
	f.Builder.WithLogger(loggerInterface)
	f.Builder.WithHttpClient(pClient)

	return f.Builder.GetInstance()
}
