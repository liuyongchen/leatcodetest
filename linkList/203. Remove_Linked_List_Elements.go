package linkList

//
//Remove all elements from a linked list of integers that have value val.
//
//Example:
//
//Input:  1->2->6->3->4->5->6, val = 6
//Output: 1->2->3->4->5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-linked-list-elements
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//添加哨兵结点，解决head.next 为要删除元素的情况
//移动游标时，需要参考两种情况
//当结点值为要删除的值时，前驱结点的next指向当前结点的next结点，不用移动前驱结点
//当结点值不是要删除的值时，前驱结点移动到当前结点，
//无论是否移动前驱结点，都要移动当前结点到下一结点。

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	temp := &ListNode{}
	temp.Next = head
	pre, cur := temp, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return temp.Next
}
