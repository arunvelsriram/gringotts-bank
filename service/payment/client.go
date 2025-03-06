package payment

import (
	"context"
	"gringotts-bank/pkg/http"
)

const baseUrl = "http://localhost:8083"

type Client struct {
	httpClient http.Client
}

func (c Client) GetCustomerTransactions(ctx context.Context, customerId string, transactions any) error {
	if err := c.httpClient.GetJson(ctx, baseUrl+"/customers/"+customerId+"/transactions", transactions); err != nil {
		return err
	}

	return nil
}

func NewClient(httpClient http.Client) Client {
	return Client{httpClient: httpClient}
}
