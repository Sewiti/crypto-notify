package coinlore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://api.coinlore.net/api"

// Client is a coinlore API client
type Client interface {
	GetCoin(context.Context, int) (Coin, error)
}

type client struct {
	httpClient *http.Client
}

// NewClient returns a new coinlore API client
func NewClient(timeout time.Duration) Client {
	return &client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *client) sendRequest(req *http.Request, val interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		// Decode error JSON?
		return fmt.Errorf("request %s: received %d status code", req.URL, res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(&val)
}
