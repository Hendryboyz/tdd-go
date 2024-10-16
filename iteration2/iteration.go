package iteration

import "strings"

func Repeat(character string, times int) string {
	// repeated := ""
	// var repeated string

	return repeatWithStrings(character, times)
}

func repeatWithFor(character string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += character
	}
	return
}

func repeatWithRange(character string, times int) (repeated string) {
	for range times {
		repeated += character
	}

	return
}

func repeatWithWhile(character string, times int) (repeated string) {
	for len(repeated) != times {
		repeated += character
	}
	return
}

func repeatWithStrings(character string, times int) string {
	return strings.Repeat(character, times)
}
