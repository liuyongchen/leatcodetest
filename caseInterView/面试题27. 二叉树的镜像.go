package caseInterView

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = mirrorTree(root.Left), mirrorTree(root.Right)
	return root
}

//同主站226题 Invert Binary Tree
