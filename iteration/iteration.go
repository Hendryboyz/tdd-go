package iteration

func Repeat(origin string, time int) string {
	var repeated string
	for len(repeated) != (len(origin) * time) {
		repeated += origin
	}
	return repeated
}
