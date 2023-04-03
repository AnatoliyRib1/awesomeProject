package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	var k = 3
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	nums2 := make([]int, 0)
	nums2 = append(nums2, nums[(len(nums)-k):]...)
	nums2 = append(nums2, nums[:(len(nums)-k)]...)
	copy(nums, nums2)

	fmt.Println(nums)
}
