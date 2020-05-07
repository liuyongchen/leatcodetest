package mathtest

//012012201
//000111222

func sortColors(nums []int) {
	a := 0
	b := len(nums) - 1
	for c := 0; c <= b; {
		if nums[c] == 0 {
			nums[c], nums[a] = nums[a], nums[c]
			a++
			c++
		} else if nums[c] == 2 {
			nums[c], nums[b] = nums[b], nums[c]
			b--
		} else {
			c++
		}
	}
}
