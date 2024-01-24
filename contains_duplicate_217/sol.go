package contains_duplicate_217

func containsDuplicate(nums []int) bool {
	hash := make(map[int]struct{})
	for _, num := range nums {
		if _, hasNum := hash[num]; hasNum {
			return true
		}
		hash[num] = struct{}{}
	}
	return false
}
