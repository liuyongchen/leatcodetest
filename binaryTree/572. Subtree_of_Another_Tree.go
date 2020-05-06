package binaryTree

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
func check(s, t *TreeNode) bool {
	if t == nil && s == nil {
		return true
	}
	if t == nil || s == nil {
		return false
	}
	if s.Val == t.Val {
		return check(s.Left, t.Left) && check(s.Right, t.Right)
	}
	return false
}
