package mathtest

import (
	"math"
)

//
//内置方法
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

//二分查找
func mySqrt0(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x>>1
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

// 作者：qu-zhen12
// 链接：https://leetcode-cn.com/problems/sqrtx/solution/er-fen-cha-zhao-he-niu-dun-die-dai-fa-by-qu-zhen12/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//好理解的
func mySqrt1(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/sqrtx/solution/x-de-ping-fang-gen-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
