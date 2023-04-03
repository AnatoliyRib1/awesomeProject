package middle_node

import (
	. "sample/utils/linkedlist"
)

func middleNode(head *ListNode) *ListNode {
	sum := 0
	start := head
	if head == nil {
		return head
	}
	for head.Next != nil {
		sum += 1
		head = head.Next
	}
	for x := sum; x > sum/2; x-- {
		start = start.Next
	}
	return start
}
