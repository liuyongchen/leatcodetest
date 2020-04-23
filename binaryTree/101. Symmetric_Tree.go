package binaryTree

//1递归解决方法，除了关心节点分支还要关心节点值
func isSymmetric1(root *TreeNode)bool{
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

//2 迭代
func isSymmetric2(root *TreeNode)bool{
	// if root == nil{
	// 	return true
	// }
	lqueue := []*TreeNode{}
	rqueue := []*TreeNode{}
	lqueue = append(lqueue,root)
	rqueue = append(rqueue,root)
	// lqueue = append(lqueue,root.Right)
	// rqueue = append(rqueue,root.Left)
	for len(lqueue)!=0 && len(rqueue)!=0{
		lcur,rcur := lqueue[0],rqueue[0]
		lqueue,rqueue = lqueue[1:],rqueue[1:]
		if lcur==nil && rcur==nil{
			continue
		}else if lcur!=nil && rcur!=nil && lcur.Val==rcur.Val{
			lqueue = append(lqueue,lcur.Left,lcur.Right)
			rqueue = append(rqueue,rcur.Right,rcur.Left)

		}else {return false}

	}
	if len(lqueue)==0 && len(rqueue)==0{
		return true
	}else {return false}
}


//3 管道