package mergeTwoLists

import . "sample/utils/linkedlist"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{}
	cur := dummy

	for list2 != nil && list1 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list2 == nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}
