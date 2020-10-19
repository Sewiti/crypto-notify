package coinlore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.coinlore.net/api"

// Coin information
type Coin struct {
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

// GetCoin information
func GetCoin(ctx context.Context, id string) (coin Coin, err error) {
	url := fmt.Sprintf(baseURL+"/ticker/?id=%s", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return coin, err
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return coin, err
	}

	err = json.NewDecoder(res.Body).Decode(&coin)
	return coin, err
}
