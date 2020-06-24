package mathtest

import (
	"fmt"
	"sort"
)

//未完成，无法参考
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	lenth := len(nums)
	start := 0
	if lenth < 3 {
		return nil
	}
	var res = [][]int{}
	for start < lenth && nums[start] <= 0 {
		if start > 0 && nums[start] == nums[start-1] {
			start++
			continue
		}
		left := start + 1
		right := lenth - 1
		for left < right {
			new := []int{}
			if left > start+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < lenth-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[start]+nums[left]+nums[right] > 0 {
				right--
				continue
			} else if nums[start]+nums[left]+nums[right] < 0 {
				left++
				continue
			}
			if left == right {
				continue
			}
			if nums[start]+nums[left]+nums[right] == 0 {
				new = append(new, nums[start], nums[left], nums[right])
				res = append(res, new)
			}
			left++
		}
		start++
	}
	return res
}
