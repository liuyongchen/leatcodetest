package strtest

import (
	"strings"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	m := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	stack := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			val, ok := m[string(s[i])]
			if ok {
				top := stack[len(stack)-1]
				if val == top {
					stack = stack[:len(stack)-1]
					continue
				}
			}
		}
		stack = append(stack, string(s[i]))

	}
	return len(stack) == 0
}

//性能太低了
func isValid0(s string) bool {
	for {
		old := s
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		if s == "" {
			return true
		}
		if s == old {
			return false
		}
	}
	return false
}

//作者：dan-bai-xi
//链接：https://leetcode-cn.com/problems/valid-parentheses/solution/jian-dan-cu-bao-jiang-fu-he-de-gua-hao-yi-ge-ge-qu/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//主站最小内存范例
func isValid2(s string) bool {
	//测试空字符串
	if s == "" {
		return true
	}

	//定义一个栈
	stack := make([]rune, len(s))
	stackLength := 0

	for _, v := range s {
		if v == '(' ||
			v == '[' ||
			v == '{' {
			//左括号,入栈
			stack[stackLength] = v
			stackLength++
		} else {
			//右括号,比较栈顶,匹配则移除,不匹配return false
			if stackLength == 0 {
				return false
			}
			if (v == ')' && stack[stackLength-1] == '(') ||
				(v == ']' && stack[stackLength-1] == '[') ||
				(v == '}' && stack[stackLength-1] == '{') {
				stackLength--
			} else {
				return false
			}
		}
	}

	return stackLength == 0
}
