package mathtest

//双指针
func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1

	for low < high {
		if target == numbers[low]+numbers[high] {
			return []int{low + 1, high + 1}
		}
		if target > numbers[low]+numbers[high] {
			high--
		} else {
			low++
		}
	}
	return []int{}
}

//二分法
func twoSum0(numbers []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(numbers)-1; i++ {
		r := search(numbers, i+1, len(numbers)-1, target-numbers[i])
		if r != -1 {
			result[0] = i + 1
			result[1] = r + 1
			break
		}
	}
	return result
}

func search(numbers []int, l, r, target int) int {
	if l > r {
		return -1
	}
	//取中值方法
	mid := l + (r-l)/2
	if numbers[mid] == target {
		return mid
	} else if numbers[mid] < target {
		return search(numbers, mid+1, r, target)
	} else {
		return search(numbers, l, mid-1, target)
	}
}
