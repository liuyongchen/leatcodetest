package caseInterView

func singleNumbers(nums []int) []int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = struct{}{}
		}
	}
	res := []int{}
	for k := range m {
		res = append(res, k)
	}
	return res
}

//位运算，需要学习
