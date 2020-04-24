package linkList

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	if head == nil{
		return nil
	}
	for head != nil{
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}
