package linkList

import (
	"sort"
)

//没思路，题解看不懂，多数用快排

//链表转切片-切片转链表
func sortList(head *ListNode) *ListNode {
	nums := make([]int, 0)
	temp := head
	for temp != nil {
		nums = append(nums, temp.Val)
		temp = temp.Next
	}
	sort.Ints(nums)
	newHead := &ListNode{}
	node := newHead
	for i := 0; i < len(nums); i++ {
		temp := &ListNode{nums[i], nil}
		node.Next = temp
		node = node.Next
	}
	return newHead.Next
}

//快排没思路
func sortList0(head *ListNode) *ListNode {
	return nil

}
