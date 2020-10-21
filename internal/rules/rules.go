package rules

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Rules set that triggers based on crypto price
type Rules []Rule

// Rule that triggers based on crypto price
type Rule struct {
	CryptoID  int     `json:"crypto_id"` // Cryptocurrency ID
	Price     float64 `json:"price"`     // Price to which rule is compared to in order to trigger
	Operator  string  `json:"rule"`      // Operator by which rule is compared to the price in order to trigger
	Triggered bool    `json:"triggered"` // Was rule triggered
}

// Read rules from a JSON file
func Read(name string) (Rules, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rules := Rules{}
	err = json.NewDecoder(file).Decode(&rules)
	if err != nil {
		return nil, err
	}

	return rules, err
}

// Write rules to a JSON file
func Write(name string, r Rules) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(r)
}

// Check checks if price triggers this rule and updates it accordingly. Returns an error if rule was already triggered.
func (r *Rule) Check(price float64) (bool, error) {
	if r.Triggered {
		return false, fmt.Errorf("check: rule has already been triggered")
	}

	var trig bool

	switch strings.ToLower(r.Operator) {
	case "lt":
		trig = price < r.Price

	case "le":
		trig = price <= r.Price

	case "gt":
		trig = price > r.Price

	case "ge":
		trig = price >= r.Price

	case "eq":
		trig = price == r.Price

	case "ne":
		trig = price != r.Price

	default:
		return false, fmt.Errorf("check %s: invalid operator", r.Operator)
	}

	r.Triggered = trig

	return trig, nil
}
