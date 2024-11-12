package romannumber

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumeral = []RomanNumeral{
	{Value: 1000, Symbol: "M"},
	{Value: 900, Symbol: "CM"},
	{Value: 500, Symbol: "D"},
	{Value: 400, Symbol: "CD"},
	{Value: 100, Symbol: "C"},
	{Value: 90, Symbol: "XC"},
	{Value: 50, Symbol: "L"},
	{Value: 40, Symbol: "XL"},
	{Value: 10, Symbol: "X"},
	{Value: 9, Symbol: "IX"},
	{Value: 5, Symbol: "V"},
	{Value: 4, Symbol: "IV"},
	{Value: 1, Symbol: "I"},
}

func ConvertRoman(arabic int) string {

	var romanNumber strings.Builder
	for _, roman := range allRomanNumeral {
		for arabic >= roman.Value {
			romanNumber.WriteString(roman.Symbol)
			arabic -= roman.Value
		}
	}
	return romanNumber.String()
}

func ConvertArabic(roman string) int {
	var arabic int
	for _, numeral := range allRomanNumeral {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return arabic
}
