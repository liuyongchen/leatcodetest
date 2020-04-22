package main

import (
	"strconv"
	"strings"
	"time"
)

//Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD,
//return the day number of the year.

//1、time标准库解法
func DayOfYear(date string) int {
	const format ="2006-01-02"
	t, _ := time.Parse(format, date)
	return t.YearDay()
}

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	start := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	return int(t.Unix() - start.Unix()) / (3600*24) + 1 //+1补上日期当天
}

//2、切片取日期法,好理解好像到，但不如time标准库简短。
// 虽然他可以让4整除，但是整数（后面有2个0以上的）的年份要用400来整除。
// 四年一闰；百年不闰,四百年再闰；千年不闰,四千年再闰；万年不闰，五十万年再闰
// 公历纪年法中，能被4整除的大多是闰年，能被100整除而不能被400整除的年份不是闰年，
// 能被3200整除的也不是闰年，如1900年是平年，2000年是闰年，3200年不是闰年。
func DayOfYear1(date string) int {
	var days = []int{0,31,28,31,30,31,30,31,31,30,31,30,31}
	dateSli := strings.Split(date,"-")
	year,_ := strconv.Atoi(dateSli[0])
	month,_ := strconv.Atoi(dateSli[1])
	day,_ := strconv.Atoi(dateSli[2])
	if year % 100 == 0{
		if year % 400 == 0 {
			days[2]=29
		}
	}else if year%4==0{
		days[2]=29
	}
	var result int
	for i:=0;i<month;i++ {
		result += days[i]
	}
	return result+day
}