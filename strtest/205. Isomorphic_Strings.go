package strtest

import (
	"bytes"
)

// 给定两个字符串 s 和 t，判断它们是否是同构的。
//
// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//
// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
//
// 示例 1:
//
// 输入: s = "egg", t = "add"
// 输出: true
// 示例 2:
//
// 输入: s = "foo", t = "bar"
// 输出: false
// 示例 3:
//
// 输入: s = "paper", t = "title"
// 输出: true
// 说明:
// 你可以假设 s 和 t 具有相同的长度。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/isomorphic-strings
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//不理解，
func isIsomorphic(s string, t string) bool {

		sb := []byte(s)
		st := []byte(t)

		if len(sb) == 0 && len(st) == 0 {
		return true
	}

		if len(sb) != len(st) {
		return false
	}

		for i := 0; i < len(sb); i++ {
		if bytes.IndexByte(sb, sb[i]) != bytes.IndexByte(st, st[i]) {
		return false
	}
	}

		return true
	}

	// 作者：liangran2019
	// 链接：https://leetcode-cn.com/problems/isomorphic-strings/solution/go-by-liangran2019/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。