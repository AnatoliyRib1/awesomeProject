package rotateArray

func rotate(nums []int, k int) []int {
	k %= len(nums)
	if len(nums) != 1 {
		nums2 := make([]int, 0)
		nums2 = append(nums2, nums[(len(nums)-k):]...)
		nums2 = append(nums2, nums[:(len(nums)-k)]...)
		copy(nums, nums2)

	}
	return nums
}
