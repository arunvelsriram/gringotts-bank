package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Client struct {
	httpClient *http.Client
}

func (c Client) GetJson(ctx context.Context, url string, response any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("unexpected status code: %d, failed to read body: %v", res.StatusCode, err)
		}
		return fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, string(body))
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
