package delete_duplicates

import . "sample/utils/linkedlist"

func deleteDuplicates(head *ListNode) *ListNode {

	start := head
	if head == nil {
		return head
	}
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return start
}
