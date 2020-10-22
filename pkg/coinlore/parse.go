package coinlore

import (
	"strconv"
)

// parse api response with proper numeric variables, zero values if string is empty
func parse(rc rawCoin) (Coin, error) {
	c := Coin{}
	var err error

	c.ID, err = parseInt(rc.ID)
	if err != nil {
		return Coin{}, err
	}

	c.Symbol = rc.Symbol
	c.Name = rc.Name
	c.NameID = rc.NameID
	c.Rank = rc.Rank

	c.PriceUSD, err = parseFloat64(rc.PriceUSD)
	if err != nil {
		return Coin{}, err
	}

	c.PercentChange24h, err = parseFloat64(rc.PercentChange24h)
	if err != nil {
		return Coin{}, err
	}

	c.PercentChange1h, err = parseFloat64(rc.PercentChange1h)
	if err != nil {
		return Coin{}, err
	}

	c.PercentChange7d, err = parseFloat64(rc.PercentChange7d)
	if err != nil {
		return Coin{}, err
	}

	c.MarketCapUSD, err = parseFloat64(rc.MarketCapUSD)
	if err != nil {
		return Coin{}, err
	}

	c.Volume24, err = parseFloat64(rc.Volume24)
	if err != nil {
		return Coin{}, err
	}

	c.Volume24Native, err = parseFloat64(rc.Volume24Native)
	if err != nil {
		return Coin{}, err
	}

	c.CSupply, err = parseFloat64(rc.CSupply)
	if err != nil {
		return Coin{}, err
	}

	c.PriceBTC, err = parseFloat64(rc.PriceBTC)
	if err != nil {
		return Coin{}, err
	}

	c.TSupply, err = parseInt64(rc.TSupply)
	if err != nil {
		return Coin{}, err
	}

	c.MSupply, err = parseInt64(rc.MSupply)
	if err != nil {
		return Coin{}, err
	}

	return c, nil
}

func parseInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	v, err := strconv.ParseInt(str, 10, 0)

	return int(v), err
}

func parseInt64(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}

	return strconv.ParseInt(str, 10, 64)
}

func parseFloat64(str string) (float64, error) {
	if str == "" {
		return 0, nil
	}

	return strconv.ParseFloat(str, 64)
}
