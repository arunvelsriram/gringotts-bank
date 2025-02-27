package http

import (
	"context"
	"encoding/json"
	"fmt"
	"gringotts-bank/pkg/log"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
)

type Client struct {
	httpClient *http.Client
}

func (c Client) GetJson(ctx context.Context, url string, response any) error {
	logger := log.Logger(ctx)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	logger.Info("http request completed", zap.String("url", url), zap.String("status", res.Status))

	if res.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			logger.Error("failed to read body", zap.String("url", url), zap.String("status", res.Status), zap.Error(err))
			return fmt.Errorf("failed to read body")
		}

		logger.Error("unexpected http status from server", zap.String("status", res.Status), zap.String("url", url), zap.String("body", string(body)))
		return fmt.Errorf("unexpected status from server: %d, body: %s", res.StatusCode, string(body))
	}

	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(response)
}

func NewClient() Client {
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	return Client{
		httpClient: client,
	}
}
