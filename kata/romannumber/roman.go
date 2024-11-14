package romannumber

import "strings"

type RomanNumber struct {
	Value int
	Roman string
}

var AllRomanNumbers = []RomanNumber{
	{Value: 1000, Roman: "M"},
	{Value: 900, Roman: "CM"},
	{Value: 500, Roman: "D"},
	{Value: 400, Roman: "CD"},
	{Value: 100, Roman: "C"},
	{Value: 90, Roman: "XC"},
	{Value: 50, Roman: "L"},
	{Value: 40, Roman: "XL"},
	{Value: 10, Roman: "X"},
	{Value: 9, Roman: "IX"},
	{Value: 5, Roman: "V"},
	{Value: 4, Roman: "IV"},
	{Value: 1, Roman: "I"},
}

func ConvertRoman(arabic int) string {
	var roman strings.Builder
	for _, number := range AllRomanNumbers {
		for arabic >= number.Value {
			arabic -= number.Value
			roman.WriteString(number.Roman)
		}
	}
	return roman.String()
}

func ConvertArabic(roman string) int {
	arabic := 0
	for _, number := range AllRomanNumbers {
		for strings.HasPrefix(roman, number.Roman) {
			roman = strings.TrimPrefix(roman, number.Roman)
			arabic += number.Value
		}
	}
	return arabic
}
