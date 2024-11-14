package romannumber

import (
	"fmt"
	"testing"
)

func TestRomanNumber(t *testing.T) {
	testCases := []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 16, Roman: "XVI"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("convert %d to %s", test.Arabic, test.Roman), func(t *testing.T) {
			expected := test.Roman
			roman := ConvertRoman(test.Arabic)

			if expected != roman {
				t.Errorf("expect %q but got %q", expected, roman)
			}
		})
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("convert %s to %d", test.Roman, test.Arabic), func(t *testing.T) {
			expected := test.Arabic
			arabic := ConvertArabic(test.Roman)

			if expected != arabic {
				t.Errorf("expected %d but got %d", expected, arabic)
			}
		})
	}

}
