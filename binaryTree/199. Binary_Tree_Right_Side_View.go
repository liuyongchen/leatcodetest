package binaryTree


// Given a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	if root.Left==nil && root.Right==nil{
		return []int{root.Val}
	}else {
		leftView := rightSideView(root.Left)
		rightView := rightSideView(root.Right)
		return append([]int{root.Val},merge(leftView,rightView)...)
	}
}

func merge(left []int,right []int)[]int  {
	var result []int
	if len(left)>len(right){
		result = make([]int,len(left))
	}else {
		result = make([]int,len(right))
	}
	for i:=0;i<len(result);i++{
		if i<len(right){
			result[i] = right[i]
		}else {
			result[i] = left[i]
		}
	}
	return result
}