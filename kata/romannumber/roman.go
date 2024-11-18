package romannumber

import "strings"

type Numeral struct {
	Value  int
	Symbol string
}

var AllRomanMap = []Numeral{
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
	var roman strings.Builder
	for _, numeral := range AllRomanMap {
		for arabic >= numeral.Value {
			arabic -= numeral.Value
			roman.WriteString(numeral.Symbol)
		}
	}
	return roman.String()
}

func ConvertArabic(roman string) (arabic int) {
	for _, numeral := range AllRomanMap {
		for strings.HasPrefix(roman, numeral.Symbol) {
			roman = strings.TrimPrefix(roman, numeral.Symbol)
			arabic += numeral.Value
		}
	}
	return arabic
}
