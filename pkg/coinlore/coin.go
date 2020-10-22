package coinlore

import (
	"context"
	"fmt"
	"net/http"
)

// Coin contains available information about a cryptocurrency, returned by an API.
type Coin struct {
	ID               int
	Symbol           string
	Name             string
	NameID           string
	Rank             int
	PriceUSD         float64
	PercentChange24h float64
	PercentChange1h  float64
	PercentChange7d  float64
	MarketCapUSD     float64
	Volume24         float64
	Volume24Native   float64
	CSupply          float64
	PriceBTC         float64
	TSupply          int64
	MSupply          int64
}

// rawCoin is a raw structure that the API responds with.
type rawCoin struct {
	ID               string `json:"id"`
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	NameID           string `json:"nameid"`
	Rank             int    `json:"rank"`
	PriceUSD         string `json:"price_usd"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange7d  string `json:"percent_change_7d"`
	MarketCapUSD     string `json:"market_cap_usd"`
	Volume24         string `json:"volume24"`
	Volume24Native   string `json:"volume24_native"`
	CSupply          string `json:"csupply"`
	PriceBTC         string `json:"price_btc"`
	TSupply          string `json:"tsupply"`
	MSupply          string `json:"msupply"`
}

// GetCoin sends a request to the api for coin's information.
func (c *client) GetCoin(ctx context.Context, id int) (Coin, error) {
	url := fmt.Sprintf(baseURL+"/ticker/?id=%d", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Coin{}, err
	}

	req = req.WithContext(ctx)

	var res []rawCoin
	err = c.sendRequest(req, &res)
	if err != nil {
		return Coin{}, err
	}

	return parse(res[0]) // Why do I have to parse this again?
}
