package caseInterView

//1.常规方法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	countA, countB := 0, 0
	for curA != nil {
		curA = curA.Next
		countA++
	}
	for curB != nil {
		curB = curB.Next
		countB++
	}
	if countB > countA {
		for i := 0; i < countB-countA; i++ {
			headB = headB.Next
		}
	} else if countA > countB {
		for i := 0; i < countA-countB; i++ {
			headA = headA.Next
		}
	}

	for headA != nil {
		if headB == headA {
			return headA
		}
		headA = headA.Next
		headB = headB.Next

	}
	return nil
}

//2 map查找
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

//3 切片查找
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	nodsA, nodsB := getNode(headA), getNode(headB)
	var result *ListNode = nil
	for i := 0; i < len(nodsA); i++ {
		if i >= len(nodsB) {
			break
		}
		if nodsB[len(nodsB)-i-1] == nodsA[len(nodsA)-i-1] {
			result = nodsA[len(nodsA)-i-1]
		}
	}
	return result
}

func getNode(head *ListNode) []*ListNode {
	var result = []*ListNode{}
	if head == nil {
		return nil
	}
	for head != nil {
		result = append(result, head)
		head = head.Next
	}
	return result
}

//4 双指针两次循环
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
