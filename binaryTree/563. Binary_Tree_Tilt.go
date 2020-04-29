package binaryTree

import (
	"math"
)

var resultn int

func findTilt(root *TreeNode) int {
	result = 0
	dfsFind(root)
	return resultn
}

func dfsFind(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfsFind(root.Left)
	right := dfsFind(root.Right)
	result += int(math.Abs(float64(left - right)))

	return left + right + root.Val
}
