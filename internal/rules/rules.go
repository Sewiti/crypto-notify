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
	CryptoID  string  `json:"crypto_id"`
	Price     float64 `json:"price"`
	Operator  string  `json:"rule"`
	Triggered bool    `json:"triggered"`
}

// Read rules from a JSON file
func Read(fileName string) (rules Rules, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&rules)
	if err != nil {
		return nil, err
	}

	return rules, err
}

// Write rules to a JSON file
func Write(fileName string, rules Rules) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(rules)
}

// Check if price triggers this rule
func (r *Rule) Check(price float64) (bool, error) {
	switch strings.ToLower(r.Operator) {
	case "lt":
		return price < r.Price, nil

	case "le":
		return price <= r.Price, nil

	case "gt":
		return price > r.Price, nil

	case "ge":
		return price >= r.Price, nil

	case "eq":
		return price == r.Price, nil

	case "ne":
		return price != r.Price, nil

	default:
		return false, fmt.Errorf("check %s: invalid operator", r.Operator)
	}
}
