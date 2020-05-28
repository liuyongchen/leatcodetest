package mathtest

import (
	"fmt"
	"sort"
)

//未完成，无法参考
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	lenth := len(nums)
	start := 0
	var res = [][]int{}
	for nums[start] <= 0 {
		left := start + 1
		right := lenth - 1
		fmt.Println("第一层")
		for left < right {
			fmt.Println("第二层")
			new := []int{}
			if nums[start]+nums[left]+nums[right] > 0 {
				fmt.Println("right", right)
				right--
				continue
			} else if nums[start]+nums[left]+nums[right] < 0 {
				fmt.Println("left", left)
				left++
				continue
			}
			new = append(new, nums[start], nums[left], nums[right])
			fmt.Println(new)
			res = append(res, new)
			fmt.Println(res)
			left++

		}
		fmt.Println("start", start)
		start++
	}
	return res
}
