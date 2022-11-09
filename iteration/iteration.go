package iteration

import "strings"

func Repeat(origin string, time int) string {
	return repeatWithLibrary(origin, time)
}

func doRepeat(origin string, time int) string {
	var repeated string
	for len(repeated) != (len(origin) * time) {
		repeated += origin
	}
	return repeated
}

func repeatWithLibrary(origin string, time int) string {
	return strings.Repeat(origin, time)
}
