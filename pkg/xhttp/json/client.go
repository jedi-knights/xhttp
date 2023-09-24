package json

import (
	"encoding/json"
	"github.com/jedi-knights/xhttp/pkg/xhttp/core"
	"go.uber.org/zap"
)

// Client is a wrapper around the core client.
type Client struct {
	CoreClient *core.Client
}

// NewClient returns a new client with a nested core client.
func NewClient() *Client {
	pFactory := core.NewFactory()
	pLogger, _ := zap.NewProduction()
	pCoreClient := pFactory.CreateWithLogger(pLogger)

	return &Client{
		CoreClient: pCoreClient,
	}
}

// Get is a wrapper around the core client's Get method. It will return the status code and the body of the response.
func (c *Client) Get(requestUrl string, out interface{}) (int, error) {
	var statusCode int
	var err error
	var bytes []byte

	if statusCode, bytes, err = c.CoreClient.Get(requestUrl); err != nil {
		return statusCode, err
	}

	if err = json.Unmarshal(bytes, out); err != nil {
		return statusCode, err
	}

	return statusCode, nil
}
