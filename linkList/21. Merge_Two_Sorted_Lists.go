package linkList

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
//
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 解题思路：
// 纪要：要对自己写的代码有信心，先自己写，再想办法改bug；
// “合并两个链表，需一个哨兵结点”，来自程序员小浩，
// 利用哨兵结点，将哨兵结点的下一结点永远指向 两个连个链表中值最小的结点；
// 当一个链表指向nil 的时候，另一个链表余下全部，都加入到哨兵结点的尾部即可完成题目要求。

func mergeTwoLists0(l1 *ListNode, l2 *ListNode) *ListNode {
	nl := &ListNode{}
	result := nl
	for l1 != nil{
		if l2 == nil{
			nl.Next = l1
			return result.Next
		}
		if l1.Val < l2.Val{
			nl.Next = l1
			l1 = l1.Next
			nl = nl.Next
		}else if l2.Val <= l1.Val{
			nl.Next = l2
			l2 = l2.Next
			nl = nl.Next
		}

	}
	nl.Next = l2
	return result.Next
}

