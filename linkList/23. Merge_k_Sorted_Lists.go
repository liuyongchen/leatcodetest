package linkList

//转化成合并两个链表
func mergeKLists(lists []*ListNode) *ListNode {
	result := &ListNode{}
	for _, list := range lists {
		result.Next = mergeTwoLists(result.Next, list)
	}
	return result.Next
}

func mergeTwoLists(t1, t2 *ListNode) *ListNode {
	temp := &ListNode{}
	result := temp
	for t1 != nil && t2 != nil {
		if t1.Val <= t2.Val {
			temp.Next = t1
			t1 = t1.Next
			temp = temp.Next
		} else if t2.Val < t1.Val {
			temp.Next = t2
			t2 = t2.Next
			temp = temp.Next
		}
	}
	if t1 == nil {
		temp.Next = t2
	}
	if t2 == nil {
		temp.Next = t1
	}
	return result.Next
}

//分治法，递归思想
func mergeKLists1(lists []*ListNode) *ListNode {

	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	default:
		mid := len(lists) / 2
		return mergeTwoLists(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
	}
}
