package main

func removeOuterParentheses(S string) string {
	res := ""
	num := 0
	start, end := 0, 0
	flag := false
	for k, v := range S {
		if string(v) == "(" {
			num++
			if !flag {
				start = k
				flag = true
			}
		} else {
			num--
			if num == 0 {
				end = k
				flag = false
				res += S[start+1 : end]
				start = end
			}
		}
	}
	return res
}

//未出解
func removeOuterParentheses1(S string) string {
	res := ""
	stack := []int32{}
	// num := 0
	start := 0
	for k, v := range S {
		if v == '(' {
			res += string(stack[start+1 : k])
			stack = stack[:k]

		}
		stack = append(stack, v)
	}

	return res
}
