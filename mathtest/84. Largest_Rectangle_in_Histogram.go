package mathtest

import (
	"fmt"
	"math"
)

//暴力法
func LargestRectangleArea(heights []int) int {
	//var stackList = []int{}
	area := 0
	for i := 0; i < len(heights); i++ {
		minh := math.MaxInt32
		for j := i + 1; j < len(heights); j++ {
			minh = min(minh, heights[j])
			area = maxl(area, minh*(j-i))
			fmt.Println(area)
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxl(a, b int) int {
	if a > b {
		return a
	}
	return b
}
