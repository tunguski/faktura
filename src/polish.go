package main

import (
	"github.com/shopspring/decimal"
	"github.com/moul/number-to-words"
)

func valToPolishText(val Val) string {
	result := ""

	remainder := val.v.Shift(2).Mod(decimal.RequireFromString("100")).Round(0)
	decimal := val.v.Round(0)

	result += ntw.IntegerToPlPl(int(decimal.IntPart()))
	result += " PLN "
	result += remainder.String()
	result += "/100"

	return result
}

