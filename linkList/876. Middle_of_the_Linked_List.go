package linkList

//两次遍历
func middleNode(head *ListNode) *ListNode {
	cur := head
	var count int
	for cur != nil {
		cur = cur.Next
		count++
	}
	var half int
	for head != nil {
		if half == count/2 {
			return head
		} // }else if half == count/2 && count%2 != 0{
		//     return head.Next
		// }
		head = head.Next
		half++

	}
	return head
}

//快慢指针
func middleNode1(head *ListNode) *ListNode {
	fast := head
	mid := head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			mid = mid.Next
		}
	}
	return mid
}
