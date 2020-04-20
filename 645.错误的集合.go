package main

// 集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，
// 导致集合丢失了一个整数并且有一个元素重复。
// 给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
//
// 示例 1:
//
// 输入: nums = [1,2,2,4]
// 输出: [2,3]
// 注意:
//
// 给定数组的长度范围是 [2, 10000]。
// 给定的数组是无序的。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/set-mismatch
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findErrorNums(nums []int)[]int{
	type void struct {}
	var value void
	set := make(map[int]void)

	sameNum := 0
	lenth := len(nums)
	sum := 0

	for i:=0;i<lenth;i++{
		if _,ok := set[nums[i]];ok{
			sameNum = nums[i]
		}else {
			set[nums[i]]=value
		}
		sum+=nums[i]
	}
	//无序递增1的结合的值的和：（1+lenth）*lenth/2
	return []int{sameNum,((1+lenth)*lenth/2)-(sum-sameNum)}
}

func FindErrorNums2(nums []int)[]int  {
	lenth := len(nums)
	templist := make([]int,lenth+1)

	res1,res2 := 0,0
	for i:=0;i<lenth;i++{
		templist[nums[i]]++
	}


	for i:=1;i<lenth+1;i++{
		if templist[i]==2{
			res1 = i
		}else if templist[i] == 0 {
			res2 = i
		}
	}
	return []int{res1,res2}
}