package caseInterView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric0(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return chk(root.Left, root.Right)
}

func chk(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	leftR := chk(left.Left, right.Right)
	rightR := chk(left.Right, right.Left)
	return leftR && rightR && left.Val == right.Val
}

//迭代法
func isSymmetric1(root *TreeNode) bool {
	lqueue := []*TreeNode{}
	rqueue := []*TreeNode{}
	lqueue = append(lqueue, root)
	rqueue = append(rqueue, root)
	for len(lqueue) != 0 && len(rqueue) != 0 {
		lcur, rcur := lqueue[0], rqueue[0]
		lqueue, rqueue = lqueue[1:], rqueue[1:]
		if lcur == nil && rcur == nil {
			continue
		} else if lcur != nil && rcur != nil && lcur.Val == rcur.Val {
			lqueue = append(lqueue, lcur.Left, lcur.Right)
			rqueue = append(rqueue, rcur.Right, rcur.Left)
		} else {
			return false
		}
	}
	if len(lqueue) == 0 && len(rqueue) == 0 {
		return true
	} else {
		return false
	}
}
