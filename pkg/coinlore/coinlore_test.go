package coinlore

import (
	"context"
	"testing"
	"time"
)

func TestGetCoin(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cl := NewClient(30 * time.Second)
	c, err := cl.GetCoin(ctx, 90) // Bitcoin
	if err != nil {
		t.Fatal(err)
	}

	if c.ID != 90 ||
		c.Symbol != "BTC" ||
		c.Name != "Bitcoin" ||
		c.NameID != "bitcoin" ||
		c.PriceBTC != 1 {
		t.Fatal("Unexpected values")
	}
}

func TestGetCoinContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	cl := NewClient(30 * time.Second)
	_, err := cl.GetCoin(ctx, 90) // Bitcoin
	if err == nil {
		t.Fatal("Expected error due to cancelled context")
	}
}
