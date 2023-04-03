package remove_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	last := head
	dummy := &ListNode{Next: head}
	start := dummy

	for last != nil {
		if val == last.Val {
			start.Next = last.Next
		} else {
			start = last
		}
		last = last.Next
	}
	return dummy.Next
}
