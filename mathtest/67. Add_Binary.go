package mathtest

import (
	"strconv"
)

func AddBinary(a string, b string) string {
	lenthA, lenthB := len(a), len(b)
	n := max(lenthA, lenthB)
	c := 0
	ans := ""
	for i := 0; i < n; i++ {
		if i < lenthA {
			c += int(a[lenthA-i-1] - '0')
		}
		if i < lenthB {
			c += int(b[lenthB-i-1] - '0')
		}
		ans = strconv.Itoa(c%2) + ans
		c /= 2
	}
	if c > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
