package recommendation

import (
	"context"
	"gringotts-bank/pkg/http"
)

const baseUrl = "http://localhost:8081"

type Client struct {
	httpClient http.Client
}

func (c Client) GetRecommendations(ctx context.Context, customerId string, recommendations any) error {
	if err := c.httpClient.GetJson(ctx, baseUrl+"/customers/"+customerId+"/recommendations", recommendations); err != nil {
		return err
	}

	return nil
}

func NewClient(httpClient http.Client) Client {
	return Client{httpClient: httpClient}
}
