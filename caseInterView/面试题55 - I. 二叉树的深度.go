package caseInterView

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := maxDepth(root.Right)
	l := maxDepth(root.Left)
	if r > l {
		return r + 1
	}
	return l + 1
}

//和主站104题相同
