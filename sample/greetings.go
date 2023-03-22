package greetings

func RemoveDuplicates(nums []int) int {
	sumValid := len(nums)
	var expectedNums = make([]int, len(nums), cap(nums))
	expectedNums[0] = nums[0]
	for i := 0; i < len(expectedNums); i++ {
		for j := i + 1; j < len(expectedNums); j++ {
			if expectedNums[i] == expectedNums[j] {
				sumValid--
				expectedNums = expectedNums[:j]
			} else {
				expectedNums = append(expectedNums, nums[j])
			}

		}

	}
	return sumValid
}
