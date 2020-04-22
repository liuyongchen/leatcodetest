package main

import (
	"strings"
	"unicode"
)

//1、双指针思路
func isPalindrome(s string) bool {
	// if len(s)==0{
	// 	 return true
	// }
	s = strings.ToLower(s)
	left,right := 0,len(s)-1
	for left < right {
		if !('a'<=s[left] && s[left]<='z' || '0'<=s[left] && s[left]<='9'){
			left++
			continue
		}
		if !('a'<=s[right] && s[right]<='z' || '0'<=s[right] && s[right]<='9'){
			right++
			continue
		}
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}

//执行用时得分低于直接处理字符
//2、标准库替换手动判断
func IsPalindrome1(s string) bool {
	// if len(s)==0{
	// 	return true
	// }
	s = strings.ToLower(s)
	left,right := 0,len(s)-1
	for left < right {
		if !(unicode.IsDigit(rune(s[left])) || unicode.IsLetter(rune(s[left]))){
			left++
			continue
		}
		if !(unicode.IsDigit(rune(s[right])) || unicode.IsLetter(rune(s[right]))){

			right--
			continue
		}
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}