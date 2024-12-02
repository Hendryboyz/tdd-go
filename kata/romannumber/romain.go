package romannumber

import "strings"

var arabicRomanMap = []struct {
	symbol string
	arabic int
}{
	{symbol: "M", arabic: 1000},
	{symbol: "CM", arabic: 900},
	{symbol: "D", arabic: 500},
	{symbol: "CD", arabic: 400},
	{symbol: "C", arabic: 100},
	{symbol: "XC", arabic: 90},
	{symbol: "L", arabic: 50},
	{symbol: "XL", arabic: 40},
	{symbol: "X", arabic: 10},
	{symbol: "IX", arabic: 9},
	{symbol: "V", arabic: 5},
	{symbol: "IV", arabic: 4},
	{symbol: "I", arabic: 1},
}

func ConvertRoman(arabic int) string {
	var roman strings.Builder
	for _, mapping := range arabicRomanMap {
		for arabic >= mapping.arabic {
			arabic -= mapping.arabic
			roman.WriteString(mapping.symbol)
		}
	}
	return roman.String()
}
