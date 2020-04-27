package main

import (
	"sort"
)

//1内置方法解题
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

//2构造解题方法，从后往前赋值
func merge1(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m < 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
