package binaryTree

//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	serch := dfsSerch(root, -1<<63, 1<<63-1)
	return serch
}

func dfsSerch(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val > min && root.Val < max && dfsSerch(root.Left, min, root.Val) &&
		dfsSerch(root.Right, root.Val, max) {
		return true
	}
	return false
}

//自制队列
func isValidBST0(root *TreeNode) bool {
	//声明队列
	queue := []*TreeNode{root}
	minQ := []int{-1 << 63}
	maxQ := []int{1<<63 - 1}

	//开始循环遍历树
	for len(queue) > 0 {
		root, min, max := queue[0], minQ[0], maxQ[0]
		if root.Val <= min || root.Val >= max {
			return false
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
			minQ = append(minQ, min)
			maxQ = append(maxQ, root.Val)
		}
		if root.Right != nil {
			minQ = append(minQ, root.Val)
			maxQ = append(maxQ, max)
			queue = append(queue, root.Right)
		}
		queue, minQ, maxQ = queue[1:], minQ[1:], maxQ[1:]
	}
	return true

}

//中序遍历
