package binaryTree

import (
	"math"
)

//Abs方法很重要，math包需要详细了解
//返回两个float64类型的绝对值，左右子树均为平衡二叉树，整个二叉树才是平衡二叉树

func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return math.Abs(float64(getStep(root.Left) - getStep(root.Right))) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}

func getStep(root *TreeNode)int{
	var l,r int
	if root == nil{
		return 0
	}
	l = getStep(root.Left)
	r = getStep(root.Right)
    if l < r{
    	return r+1
	}
	return l+1
}