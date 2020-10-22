package coinlore

import "testing"

func TestParse(t *testing.T) {
	rc := rawCoin{
		ID:               "90",
		Symbol:           "BTC",
		Name:             "Bitcoin",
		NameID:           "bitcoin",
		Rank:             1,
		PriceUSD:         "12786.24",
		PercentChange24h: "7.24",
		PercentChange1h:  "1.05",
		PercentChange7d:  "12.02",
		MarketCapUSD:     "236488925301.19",
		Volume24:         "111557152846.57",
		Volume24Native:   "8724780.27",
		CSupply:          "18495577.00",
		PriceBTC:         "1.00",
		TSupply:          "18495577",
		MSupply:          "21000000",
	}

	c, err := parse(rc)
	if err != nil {
		t.Fatal(err)
	}

	if c.ID != 90 ||
		c.Symbol != "BTC" ||
		c.Name != "Bitcoin" ||
		c.NameID != "bitcoin" ||
		c.Rank != 1 ||
		c.PriceUSD != 12786.24 ||
		c.PercentChange24h != 7.24 ||
		c.PercentChange1h != 1.05 ||
		c.PercentChange7d != 12.02 ||
		c.MarketCapUSD != 236488925301.19 ||
		c.Volume24 != 111557152846.57 ||
		c.Volume24Native != 8724780.27 ||
		c.CSupply != 18495577.0 ||
		c.PriceBTC != 1.0 ||
		c.TSupply != 18495577 ||
		c.MSupply != 21000000 {
		t.Fatal("Unexpected values received")
	}
}

func TestParseWithEmptyValues(t *testing.T) {
	rc := rawCoin{
		ID:               "2",
		Symbol:           "DOGE",
		Name:             "Dogecoin",
		NameID:           "dogecoin",
		Rank:             35,
		PriceUSD:         "0.002646",
		PercentChange24h: "1.10",
		PercentChange1h:  "-0.17",
		PercentChange7d:  "0.46",
		MarketCapUSD:     "334065493.69",
		Volume24:         "27400214.56",
		Volume24Native:   "10354213518.91",
		CSupply:          "126239356384.00",
		PriceBTC:         "2.15E-7",
		TSupply:          "",
		MSupply:          "",
	}

	c, err := parse(rc)
	if err != nil {
		t.Fatal(err)
	}

	if c.ID != 2 ||
		c.Symbol != "DOGE" ||
		c.Name != "Dogecoin" ||
		c.NameID != "dogecoin" ||
		c.Rank != 35 ||
		c.PriceUSD != 0.002646 ||
		c.PercentChange24h != 1.10 ||
		c.PercentChange1h != -0.17 ||
		c.PercentChange7d != 0.46 ||
		c.MarketCapUSD != 334065493.69 ||
		c.Volume24 != 27400214.56 ||
		c.Volume24Native != 10354213518.91 ||
		c.CSupply != 126239356384.00 ||
		c.PriceBTC != 2.15e-7 ||
		c.TSupply != 0 ||
		c.MSupply != 0 {
		t.Fatal("Unexpected values received")
	}
}
