package containsDuplicate

func containsDuplicate(nums []int) bool {

	exists := make(map[int]bool, len(nums))
	for _, key := range nums {
		if exists[key] {
			return true
		}
		exists[key] = true
	}
	return false
}
