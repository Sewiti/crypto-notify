package main

import (
	"fmt"
	"strings"
)

func formatOp(op string) (string, error) {
	switch strings.ToLower(op) {
	case "lt":
		return "less than", nil

	case "le":
		return "less than or equals", nil

	case "gt":
		return "greater than", nil

	case "ge":
		return "greater than or equals", nil

	case "eq":
		return "equals", nil

	case "ne":
		return "not equals", nil

	default:
		return "", fmt.Errorf("format %s: invalid operator", op)
	}
}
