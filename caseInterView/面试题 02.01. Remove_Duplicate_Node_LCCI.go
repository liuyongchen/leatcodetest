package caseInterView

func removeDuplicateNodes(head *ListNode) *ListNode {
	pre := &ListNode{}
	cur := head
	pre.Next = head
	m := make(map[int]struct{})
	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		m[cur.Val] = struct{}{}
		pre = cur
		cur = cur.Next
	}
	return head
}
