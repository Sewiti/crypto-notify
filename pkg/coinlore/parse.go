package coinlore

import "strconv"

// parse api response with proper numeric variables
func parse(rc rawCoin) (Coin, error) {
	c := Coin{}

	i, err := strconv.ParseInt(rc.ID, 10, 0)
	if err != nil {
		return Coin{}, err
	}
	c.ID = int(i)

	c.Symbol = rc.Symbol
	c.Name = rc.Name
	c.NameID = rc.NameID
	c.Rank = rc.Rank

	f, err := strconv.ParseFloat(rc.PriceUSD, 64)
	if err != nil {
		return Coin{}, err
	}
	c.PriceUSD = f

	f, err = strconv.ParseFloat(rc.PercentChange24h, 64)
	if err != nil {
		return Coin{}, err
	}
	c.PercentChange24h = f

	f, err = strconv.ParseFloat(rc.PercentChange1h, 64)
	if err != nil {
		return Coin{}, err
	}
	c.PercentChange1h = f

	f, err = strconv.ParseFloat(rc.PercentChange7d, 64)
	if err != nil {
		return Coin{}, err
	}
	c.PercentChange7d = f

	f, err = strconv.ParseFloat(rc.MarketCapUSD, 64)
	if err != nil {
		return Coin{}, err
	}
	c.MarketCapUSD = f

	f, err = strconv.ParseFloat(rc.Volume24, 64)
	if err != nil {
		return Coin{}, err
	}
	c.Volume24 = f

	f, err = strconv.ParseFloat(rc.Volume24Native, 64)
	if err != nil {
		return Coin{}, err
	}
	c.Volume24Native = f

	f, err = strconv.ParseFloat(rc.CSupply, 64)
	if err != nil {
		return Coin{}, err
	}
	c.CSupply = f

	f, err = strconv.ParseFloat(rc.PriceBTC, 64)
	if err != nil {
		return Coin{}, err
	}
	c.PriceBTC = f

	i, err = strconv.ParseInt(rc.TSupply, 10, 0)
	if err != nil {
		return Coin{}, err
	}
	c.TSupply = i

	i, err = strconv.ParseInt(rc.MSupply, 10, 0)
	if err != nil {
		return Coin{}, err
	}
	c.MSupply = i

	return c, nil
}
