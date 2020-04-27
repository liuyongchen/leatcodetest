package caseInterView

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(getStep(root.Left)-getStep(root.Right))) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}

func getStep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getStep(root.Left)
	r := getStep(root.Right)
	if r > l {
		return r
	}
	return l + 1
}
