package linkList

//判断链表是否有环
//1、快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	for fast != nil && fast.Next != nil{
		fast = head.Next.Next
		head = head.Next
		if fast==head{
			return true
		}
	}
	return false
}


//2、哈希表
func hasCycle2(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil{
		if _ ,ok := m[head];ok{
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}