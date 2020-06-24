package main

import (
	"fmt"
)

//快慢指针
func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
		fmt.Println(fast)
	}
	return fast == 1
}

func step(n int) int {
	var result int
	for n > 0 {
		result += (n % 10) * (n % 10)
		n = n / 10
	}

	return result
}

//hash集
func isHappy1(n int) bool {
	m := make(map[int]bool)
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}
