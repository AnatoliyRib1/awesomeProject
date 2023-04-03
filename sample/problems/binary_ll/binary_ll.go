package binary_ll

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var x string

	for head != nil {
		val := head.Val
		head = head.Next
		x += strconv.Itoa(val)
	}
	res, _ := strconv.ParseInt(x, 2, 64)
	return int(res)
}
