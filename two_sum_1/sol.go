package two_sum_1

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if prevIdx, found := numMap[complement]; found {
			return []int{prevIdx, i}
		}
		numMap[num] = i
	}
	return []int{}
}
