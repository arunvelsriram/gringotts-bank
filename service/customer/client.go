package customer

import (
	"context"
	"gringotts-bank/pkg/http"
)

const baseUrl = "http://localhost:8082"

type Client struct {
	httpClient http.Client
}

func (c Client) GetCustomers(ctx context.Context, customers any) error {
	if err := c.httpClient.GetJson(ctx, baseUrl+"/customers", customers); err != nil {
		return err
	}

	return nil
}

func (c Client) GetCustomer(ctx context.Context, id string, customer any) error {
	if err := c.httpClient.GetJson(ctx, baseUrl+"/customers/"+id, customer); err != nil {
		return err
	}

	return nil
}

func NewClient(httpClient http.Client) Client {
	return Client{httpClient: httpClient}
}
