package binaryTree

var result int

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	dfsDiameter(root)
	return result
}

func dfsDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfsDiameter(root.Left)
	r := dfsDiameter(root.Right)
	path := max(l, r)
	result = max(l+r, result)
	return path + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
