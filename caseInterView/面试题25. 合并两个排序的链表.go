package caseInterView

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	pre := newHead

	for l1 != nil {
		if l2 == nil {
			pre.Next = l1
			return newHead.Next
		}
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	pre.Next = l2
	return newHead.Next
}
