package linkList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for n+1 > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next

	}
	slow.Next = slow.Next.Next
	return head
}
