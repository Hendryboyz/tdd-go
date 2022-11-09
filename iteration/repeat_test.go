package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat \"a\" 5 times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"

		assertString(t, actual, expected)
	})

	t.Run("Repeat \"b\" k times", func(t *testing.T) {
		k := 3
		actual := Repeat("b", k)
		expected := "bbb"

		assertString(t, actual, expected)
	})
}

func assertString(t testing.TB, actual, expected string) {
	if actual != expected {
		t.Errorf("Expect %q but actual %q", expected, actual)
	}
}

func ExampleRepeat() {
	cRepeated := Repeat("c", 5)
	fmt.Println(cRepeated)
	// Output: ccccc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 3)
	}
}
