package mathtest

//二分法
//内存消耗较高，需要多看看
func searchInsert(nums []int, target int) int {
	return searchInh(nums, 0, len(nums)-1, target)

}

func searchInh(nums []int, l, r, target int) int {
	if l+1 >= r {

		if nums[r] <= target {
			return r
		}
		if nums[l] >= target {
			return l
		}
		return r + 1
	}

	mid := l + (r-l)/2
	if nums[mid] == target {
		return mid
	}
	var res int
	if nums[mid] > target {
		res = searchInh(nums, l, mid-1, target)
	} else {
		res = searchInh(nums, mid+1, r, target)
	}
	return res
}

//遍历
func searchInsert0(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
	}
	return len(nums)
}
