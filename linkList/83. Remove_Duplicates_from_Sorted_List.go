package linkList

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := &ListNode{}
	res.Next = head
	first := head
	var m = make(map[int]struct{})
	for first != nil {
		if _, ok := m[first.Val]; !ok {
			m[first.Val] = struct{}{}
			first = first.Next
			res = res.Next
			continue
		}
		first = first.Next
		res.Next = first
	}

	return head
}
