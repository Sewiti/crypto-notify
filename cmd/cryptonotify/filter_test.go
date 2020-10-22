package main

import (
	"testing"

	"github.com/Sewiti/crypto-notify/internal/rules"
)

func TestFilter(t *testing.T) {
	const n = 64

	var r rules.Rules

	for i := 0; i < n; i++ {
		rule := rules.Rule{
			Triggered: i%4 == 0,
		}

		r = append(r, rule)
	}

	r = filter(r)

	for _, v := range r {
		if v.Triggered {
			t.Fatal("Filter returned slice including a triggered rule")
		}
	}
}

func TestTableDistinct(t *testing.T) {
	tests := []struct {
		input rules.Rules
	}{
		{
			rules.Rules{
				{CryptoID: 10},
				{CryptoID: 20},
				{CryptoID: 30},
				{CryptoID: 10},
				{CryptoID: 10},
				{CryptoID: 20},
				{CryptoID: 10},
			},
		},
		{
			rules.Rules{
				{CryptoID: 10},
				{CryptoID: 10},
				{CryptoID: 10},
				{CryptoID: 10},
				{CryptoID: 10},
			},
		},
		{
			rules.Rules{
				{CryptoID: 50},
				{CryptoID: 10},
				{CryptoID: 20},
				{CryptoID: 40},
				{CryptoID: 50},
				{CryptoID: 20},
				{CryptoID: 60},
			},
		},
	}

	for _, test := range tests {
		c := distinct(test.input)

		m := make(map[int]struct{})

		for _, v := range c {
			if _, exists := m[v]; exists {
				t.Fatalf("Returned slice has duplicating values: %d", v)
			} else {
				m[v] = struct{}{}
			}
		}
	}
}
