package romannumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RomanTest struct {
	Arabic int
	Roman  string
}

type RomanNumeralTestSuite struct {
	suite.Suite
	TestCases []RomanTest
}

func (s *RomanNumeralTestSuite) SetupTest() {
	s.TestCases = []RomanTest{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
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
}

func (s *RomanNumeralTestSuite) TestConvertRoman() {
	for _, eachCase := range s.TestCases {
		s.Run(fmt.Sprintf("convert %d to %s", eachCase.Arabic, eachCase.Roman), func() {
			roman := ConvertRoman(eachCase.Arabic)
			expected := eachCase.Roman
			s.Equal(roman, expected, "expected %q but got %q", expected, roman)
		})
	}
}

func (s *RomanNumeralTestSuite) TestConvertArabic() {
	for _, eachCase := range s.TestCases {
		s.Run(fmt.Sprintf("convert %s to %d", eachCase.Roman, eachCase.Arabic), func() {
			arabic := ConvertArabic(eachCase.Roman)
			expected := eachCase.Arabic
			s.Equal(arabic, expected, "expected %d but got %d", expected, arabic)
		})
	}
}

func (s *RomanNumeralTestSuite) TeardownTest() {

}

func TestRomanNumeral(t *testing.T) {
	suite.Run(t, new(RomanNumeralTestSuite))
}
