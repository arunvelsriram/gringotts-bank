package customer

import (
	"context"
	"gringotts-bank/pkg/http"
)

const baseUrl = "http://localhost:8081"

type Client struct {
	httpClient http.Client
}

func (c Client) GetCustomers(ctx context.Context) (Customers, error) {
	var customers Customers

	if err := c.httpClient.GetJson(ctx, baseUrl+"/customers", &customers); err != nil {
		return nil, err
	}

	return customers, nil
}

func NewClient(httpClient http.Client) Client {
	return Client{httpClient: httpClient}
}
