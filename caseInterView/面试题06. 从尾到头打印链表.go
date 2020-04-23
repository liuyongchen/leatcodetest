package caseInterView

import (
	"container/list"
)

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//  
//
// 示例 1：
//
// 输入：head = [1,3,2]
// 输出：[2,3,1]
//  
//
// 限制：
//
// 0 <= 链表长度 <= 10000
type ListNode struct {
	    Val int
	    Next *ListNode
	}

//1 栈操作
func reversePrint0(head *ListNode)[]int{
	if head == nil{
		return nil
	}
	res := list.New()
	for head != nil {
		res.PushFront(head.Val)
		head = head.Next

	}
	ret := []int{}
	for i := res.Front();i!=nil;i=i.Next(){
		ret = append(ret,i.Value.(int))
	}
	return ret
}

//2 链表逆序
func reversePrint1(head *ListNode)[]int  {
	if head == nil{
		return nil
	}
	var revHead *ListNode
	for head != nil{
		node := head.Next
		head.Next = revHead
		revHead = head
		head = node
	}
	var result []int
	for revHead != nil{
		result = append(result,revHead.Val)
		revHead = revHead.Next
	}
	return result
}

//3 数组逆序 不适合较大的数据量
func reversePrint2(head *ListNode)[]int  {
	if head == nil{
		return nil
	}
	var result []int
	for head != nil{
		result = append(result,head.Val)
		head = head.Next
	}

	for i,j:=0,len(result)-1;i<j;{
		result[i],result[j]=result[j],result[i]
		i++
		j--
	}
	return result
}

//4 递归法耗时太长
func reversePrint3(head *ListNode)[]int  {
	if head == nil{
		return nil
	}
		return appendDate(head)
}

func appendDate(head *ListNode)[]int  {
	if head.Next != nil{
		date := appendDate(head.Next)
		date = append(date,head.Val)
		return date
	}
	return []int{head.Val}
}