package core

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type CoreInterface interface {
	Get(requestUrl string) (int, []byte, error)
}

type Client struct {
	HttpClient *http.Client
	Logger     LoggerInterface
}

func NewClient() *Client {
	return &Client{}
}

// Get performs a GET request to the specified URL and returns the response body as a byte array.
func (c *Client) Get(requestUrl string) (int, []byte, error) {
	var err error
	var bytes []byte
	var pRequest *http.Request
	var pResponse *http.Response

	if requestUrl == "" {
		return 0, nil, fmt.Errorf("client: the specified URL is empty")
	}

	ctx := context.Background()
	if pRequest, err = http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil); err != nil {
		return 0, nil, fmt.Errorf("client: could not create request: %w", err)
	}

	c.Logger.Info("request", zap.String("url", requestUrl))

	if pResponse, err = c.HttpClient.Do(pRequest); err != nil {
		return 0, nil, fmt.Errorf("client: error making http request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.Logger.Error("error closing response body", zap.Error(err))
		}
	}(pResponse.Body)

	if bytes, err = io.ReadAll(pResponse.Body); err != nil {
		return 0, nil, fmt.Errorf("client: error reading response body: %w", err)
	}

	c.Logger.Info("status code", zap.Int("code", pResponse.StatusCode))
	c.Logger.Info("response body", zap.ByteString("body", bytes))

	return pResponse.StatusCode, bytes, nil
}
