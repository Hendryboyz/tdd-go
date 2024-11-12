package genericlist

func Sum(numbers []int) (sum int) {
	add := func(prev, cur int) int {
		return prev + cur
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	sumAll := func(prev, current []int) []int {
		return append(prev, Sum(current))
	}
	return Reduce(numbersToSum, sumAll, []int{})
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	sumTail := func(prev, current []int) []int {
		if len(current) <= 1 {
			return append(prev, 0)
		}
		tail := current[1:]
		return append(prev, Sum(tail))
	}
	return Reduce(numbersToSum, sumTail, []int{})
}
