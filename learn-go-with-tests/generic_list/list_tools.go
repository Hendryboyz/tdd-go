package genericlist

type Person struct {
	Name string
}

// the core algorithm in this generic function
func Reduce[T, K any](collections []K, fn func(prevValue T, currentValue K) T, initialValue T) T {
	var result T
	result = initialValue
	for _, elem := range collections {
		result = fn(result, elem)
	}
	return result
}

func Find[T any](collections []T, predictFn func(current T) bool) (T, bool) {
	var value T
	for _, elem := range collections {
		if predictFn(elem) {
			value = elem
			return value, true
		}
	}
	return value, false
}
