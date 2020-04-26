package binaryTree

type TreeNodei struct {
	Val   int
	Left  *TreeNodei
	Right *TreeNodei
}

//合并二叉树
func mergeTrees(t1 *TreeNodei, t2 *TreeNodei) *TreeNodei {

	return dfs(t1, t2)
}

func dfs(t1, t2 *TreeNodei) *TreeNodei {
	var t3 *TreeNodei = new(TreeNodei)
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t3.Val = t1.Val + t2.Val
	t3.Left = dfs(t1.Left, t2.Left)
	t3.Right = dfs(t1.Right, t2.Right)
	return t3
}

//leatcode范例 空间换时间，最优时间范例
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		t1.Left = mergeTrees1(t1.Left, t2.Left)
		t1.Right = mergeTrees1(t1.Right, t2.Right)
	} else if t1 == nil {
		t1 = &TreeNode{
			Val: t2.Val,
		}
		t1.Left = mergeTrees1(nil, t2.Left)
		t1.Right = mergeTrees1(nil, t2.Right)
	} else { //t2==nil
		t1.Left = mergeTrees1(t1.Left, nil)
		t1.Right = mergeTrees1(t1.Right, nil)
	}

	return t1
}
