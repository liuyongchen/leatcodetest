package strtest

func generateParenthesis(n int) []string {
	combinations := new([]string)
	generateALL(make([]byte, 2*n), 0, combinations)
	return *combinations
}

func generateALL(current []byte, pos int, result *[]string) {
	if pos == len(current) {
		if valid(current) {
			*result = append(*result, string(current))
		}
	} else {
		current[pos] = '('
		generateALL(current, pos+1, result)
		current[pos] = ')'
		generateALL(current, pos+1, result)
	}
}

func valid(current []byte) bool {
	balance := 0
	for _, v := range current {
		if v == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

//十分巧妙的解法
func generateParenthesis1(n int) []string {
	Output := new([]string)
	_generate(0, 0, n, "", Output)
	return *Output
}

func _generate(left int, right int, max int, s string, Output *[]string) {
	// 递归终止条件
	if left == right && left == max {
		*Output = append(*Output, s)
		return
	}
	// 递归主体
	if left < max {
		_generate(left+1, right, max, s+"(", Output)
	}
	if right < left {
		_generate(left, right+1, max, s+")", Output)
	}
}
