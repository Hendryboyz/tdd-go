package sum

func Sum(numbers []int) (sum int) {
	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfSums := len(numbersToSum)
	// sums := make([]int, lengthOfSums)
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) <= 1 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return
}
