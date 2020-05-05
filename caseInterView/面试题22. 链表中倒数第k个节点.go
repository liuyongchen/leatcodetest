package caseInterView

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	cur, prek := head, head
	for i := 0; i < k; i++ {
		if cur != nil {
			cur = cur.Next
		} else if cur == nil {
			if i < k-1 {
				return nil
			} else {
				return prek
			}
		}
	}
	for cur != nil {
		cur = cur.Next
		prek = prek.Next
	}
	return prek
}
