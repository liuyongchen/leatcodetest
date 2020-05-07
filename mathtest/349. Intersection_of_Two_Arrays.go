package mathtest

//map匹配
func intersection(nums1 []int, nums2 []int) []int {
	if nums2 == nil || nums1 == nil {
		return nil
	}
	result := []int{}
	m := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); i++ {
			if nums1[i] == nums2[j] {
				if _, ok := m[nums1[i]]; !ok {
					m[nums2[j]] = struct{}{}
					result = append(result, nums2[j])
				}
				continue
			}

		}
	}
	return result
}

//二分法 没思路
