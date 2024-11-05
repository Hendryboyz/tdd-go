package genericlist

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("Sum collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		assertArrayError(t, actual, expected, numbers)
	})

	t.Run("Sum collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		assertArrayError(t, actual, expected, numbers)
	})
}

func assertArrayError(t testing.TB, actual, expected int, numbers []int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %d want %d, %v", actual, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Sum 2 slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		assertSliceError(t, actual, expected)
	})

	t.Run("Sum 1 slice with lot of 1", func(t *testing.T) {
		actual := SumAll([]int{1, 1, 1, 1, 1, 1})
		expected := []int{6}

		assertSliceError(t, actual, expected)
	})
}

func assertSliceError(t testing.TB, actual, expected []int) {
	t.Helper()
	// only available in go 1.21
	// if !slices.Equal(actual, expected) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestSumTail(t *testing.T) {

	checkSums := func(t testing.TB, actual, expected []int) {
		t.Helper()
		// only available in go 1.21
		// if !slices.Equal(actual, expected) {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}
	}

	t.Run("Sum tail length 2 slice", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, actual, expected)
	})

	t.Run("Sum empty slices safely", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{1, 3, 9}, []int{2})
		expected := []int{0, 12, 0}

		checkSums(t, actual, expected)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		got := Reduce([]int{1, 2, 3}, multiply, 1)
		assert.Equal(t, 6, got)
	})
	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			var builder strings.Builder
			builder.WriteString(x)
			builder.WriteString(y)
			return builder.String()
		}
		got := Reduce([]string{"hello", " world"}, concatenate, "")
		assert.Equal(t, "hello world", got)
	})
}

func TestBadBank(t *testing.T) {
	var (
		henry = Account{Name: "henry", Balance: 100}
		mary  = Account{Name: "Mary", Balance: 75}
		amy   = Account{Name: "Amy", Balance: 200}

		transactions = []Transaction{
			NewTransaction(henry, amy, 50),
			NewTransaction(amy, mary, 200),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}
	assert.EqualValues(t, newBalanceFor(henry), 50)
	assert.EqualValues(t, newBalanceFor(amy), 50)
	assert.EqualValues(t, newBalanceFor(mary), 275)
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		firstEven, found := Find(numbers, func(current int) bool {
			return current&1 == 0
		})
		assert.True(t, found)
		assert.EqualValues(t, 2, firstEven)
	})

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Martin")
		})

		assert.True(t, found)
		assert.Equal(t, king, Person{Name: "Martin Fowler"})
	})
}
