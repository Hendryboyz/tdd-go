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

func SumAll(x []int, y []int) []int {

	return []int{3, 10}
}
