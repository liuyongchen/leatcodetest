package caseInterView

//耗时4ms，可以再优化，减少append的使用

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	list := []*TreeNode{}
	result := []int{}
	list = append(list, root)
	for len(list) > 0 {
		if list[0] == nil {
			list = list[1:]
			continue
		}
		result = append(result, list[0].Val)
		list = append(list, list[0].Left, list[0].Right)
		list = list[1:]
	}
	return result

}
