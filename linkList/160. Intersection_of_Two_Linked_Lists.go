package linkList


type ListNode struct {
	Val int
	Next *ListNode
}

//1循环两个连边取链表长度，再次循环链表找焦点，内存消耗较高
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA==nil || headB==nil{
		return nil
	}
	var countA int
	var countB int
	var curA = headA
	var curB = headB
	for curA!=nil{
		curA=curA.Next
		countA++
	}
	for curB!=nil{
		curB=curB.Next
		countB++
	}
	if countB < countA{
		for i:=0;i<countA-countB;i++{
			headA = headA.Next
		}
	}else {
		for i:=0;i<countB-countA;i++{
			headB = headB.Next
		}
	}

	for headB != nil{
		if headB == headA{
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

//2示例解，分别遍历链表A、一遍用map检查交点
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil{
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil{
		if _,ok := m[headB];ok{
			return headB
		}
		headB = headB.Next
	}
	return headB
}

//3空间换时间，用切片保存所有节点，适合较少数据量
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	nodesA,nodesB := getnode(headA),getnode(headB)
	var result *ListNode=nil
	for i:=0;i<len(nodesA);i++{
		if len(nodesB)<i{
			break
		}
		if nodesB[len(nodesB)-i] == nodesA[len(nodesA)-i]{
			result = nodesA[len(nodesA)-i]
		}
	}
	return result
}

func getnode(head *ListNode)[]*ListNode  {
	var nodes []*ListNode
	for head != nil{
		nodes = append(nodes,head)
		head = head.Next
	}
	return nodes
}