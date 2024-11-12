package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat char 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertError(t, expected, repeated)
	})

	t.Run("Repeat char 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertError(t, expected, repeated)
	})
}

func assertError(t testing.TB, expected, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated) // Output: bbbbb
}
