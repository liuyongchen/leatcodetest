package linkList

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := &ListNode{}
	temp.Next = head
	res := temp
	cur := head
	for cur != nil {
		temp.Next = cur.Next
		cur.Next = cur.Next.Next
		temp.Next.Next = cur
		temp = cur
		cur = cur.Next

	}
	return res
}
