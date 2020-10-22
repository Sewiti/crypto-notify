package main

import (
	"sort"

	"github.com/Sewiti/crypto-notify/internal/rules"
)

// filter filters out already triggered rules and returns a filtered rules set.
func filter(r rules.Rules) rules.Rules {
	filt := rules.Rules{}

	for _, v := range r {
		if !v.Triggered {
			filt = append(filt, v)
		}
	}

	return filt
}

// distinct filters out duplicate coins from rules set and returns unique coins slice.
func distinct(r rules.Rules) (coins []int) {
	var all []int // Coins

	for _, v := range r {
		all = append(all, v.CryptoID)
	}

	// There might be a little bit more efficient way
	sort.Ints(all)

	for i := range all {
		if i == 0 || all[i] != all[i-1] {
			coins = append(coins, all[i])
		}
	}

	return coins
}
