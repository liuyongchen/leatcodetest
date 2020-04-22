package main



func numberOfSubarrays(nums []int, k int) int {

	for i:=0;i<len(nums);i++{

		count := 0
		if count == k{break}
		if nums[i]%2 != 0{

		}
	}
	return 0
}
//解法，但不理解！！！

func numberOfSubarrays0(nums []int, k int) int {
	odd := make([]int, len(nums)+1)
	oddNum, prefixLen := 0, 0
	ans := 0
	for i := 0; i <= len(nums); i++ {
		prefixLen += 1 //自己也要算进去
		if i == len(nums) || nums[i]%2 == 1 {
			odd[oddNum] = prefixLen
			if oddNum >= k {
				ans += odd[oddNum-k] * odd[oddNum]
			}
			oddNum += 1
			prefixLen = 0
		}
	}
	return ans
}

