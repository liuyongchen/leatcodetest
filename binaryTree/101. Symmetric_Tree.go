package binaryTree

//1递归解决方法，除了关心节点分支还要关心节点值
func isSymmetric(root *TreeNode)bool{
	if root == nil {
		return true
	}
	return chk(root.Left,root.Right)
}


func chk(l,r *TreeNode)bool{
	if  l==nil && r==nil{
		return true
	}
	if l==nil||r==nil{
		return false
	}
	return chk(l.Left,r.Right)&&chk(l.Right,r.Left)&&l.Val==r.Val
}

