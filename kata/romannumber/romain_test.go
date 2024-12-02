package romannumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RomanArabic struct {
	Roman  string
	Arabic int
}

type RomanConversionSuite struct {
	suite.Suite
	testCases []RomanArabic
}

func (suite *RomanConversionSuite) SetupTest() {
	suite.testCases = []RomanArabic{
		{Roman: "I", Arabic: 1},
		{Roman: "II", Arabic: 2},
		{Roman: "IV", Arabic: 4},
		{Roman: "V", Arabic: 5},
		{Roman: "VIII", Arabic: 8},
		{Roman: "IX", Arabic: 9},
		{Roman: "X", Arabic: 10},
		{Roman: "XVIII", Arabic: 18},
		{Roman: "XL", Arabic: 40},
		{Roman: "XLIX", Arabic: 49},
		{Roman: "L", Arabic: 50},
		{Roman: "XC", Arabic: 90},
		{Roman: "C", Arabic: 100},
		{Roman: "CD", Arabic: 400},
		{Roman: "D", Arabic: 500},
		{Arabic: 900, Roman: "CM"},
		{Roman: "M", Arabic: 1000},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}
}

func (suite *RomanConversionSuite) TestConvertRoman() {
	for _, testCase := range suite.testCases {
		suite.Run(fmt.Sprintf("convert %s to %d", testCase.Roman, testCase.Arabic), func() {
			result := ConvertRoman(testCase.Arabic)
			suite.Equal(testCase.Roman, result)
		})
	}
}

func (s *RomanConversionSuite) TeardownTest() {}

func TestRomanConversionSuite(t *testing.T) {
	suite.Run(t, new(RomanConversionSuite))
}
