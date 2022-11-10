package sum

func Sum(nums []int) int {
	var sum int
	// for i := 0; i < len(nums); i++ {
	// 	sum += nums[i]
	// }
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for index, numbers := range numbersToSum {
		sums[index] = Sum(numbers)
	}
	return sums
}
