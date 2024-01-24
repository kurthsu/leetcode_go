package two_sum2_167

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		num := numbers[left] + numbers[right]
		if num > target {
			right--
		} else if num < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{0, 0}
}
