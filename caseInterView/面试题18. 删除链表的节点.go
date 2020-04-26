package caseInterView

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	empty := &ListNode{}
	cur := head
	pre := empty
	pre.Next = head
	for cur != nil {
		if cur.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return empty.Next
}
