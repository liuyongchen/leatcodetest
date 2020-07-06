package mathtest

import (
	"fmt"
	"sort"
)

//效率很糟糕，内存也不是很好，建议不用纠结用几个for循环
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	start, end := 0, len(nums)-1
	res := [][]int{}
	for start < end {
		fmt.Println(start, end)
		if start > 0 && nums[start] == nums[start-1] {
			if start == end-3 {
				start = 0
				end--
				continue
			} else {
				start++
				continue
			}
		}
		if end < len(nums)-1 && nums[end] == nums[end+1] {
			fmt.Println("nums[end]:", nums[end])
			end--
			continue
		}
		left, right := start+1, end-1
		for left < right {
			new := []int{}
			if left > start+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < end-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if target > nums[start]+nums[end]+nums[left]+nums[right] {
				left++
				continue
			} else if target < nums[start]+nums[end]+nums[left]+nums[right] {
				right--
				continue
			} else {
				new = append(new, nums[start], nums[left], nums[right], nums[end])
				res = append(res, new)
			}
			left++
		}
		if start == end-3 {
			start = 0
			end--
		} else {
			start++
		}
	}
	return res
}
