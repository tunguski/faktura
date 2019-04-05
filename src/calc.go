package main

import (
	"github.com/shopspring/decimal"
)

var (
	zero = (func () decimal.Decimal {
		num, _ := decimal.NewFromString("0")
		return num
	})()
)

func sum(values []string) decimal.Decimal {
	result := zero
	for _, num := range values {
		newVal, _ := decimal.NewFromString(num)
		result = result.Add(newVal)
	}
	return result
}

