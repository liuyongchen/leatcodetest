package main

import (
	"fmt"
)

type Persion struct {
	Name string
	Age int64
}

type Student struct {
	Persion Persion
	Grade string

}

func main()  {
	// newS := Student{
	// 	Persion: Persion{
	// 		Name:"petter",
	// 		Age:18,
	// 	},
	// 	Grade:   "12",
	// }
	// valueOf := reflect.ValueOf(newS)
	// valueOf.Elem()
	// for i:=0;i< valueOf.NumField();i++{
	// 	fmt.Println(valueOf.Field(i))
	// }
	a := make([]int,7)
	list1 := []int{1,2,4,6,4,5}
	nums2 := FindErrorNums2(list1)
	fmt.Println(nums2)
	for i:=0;i<len(list1);i++{
		a[list1[i]]++
	}
	fmt.Println(list1)
	fmt.Println(a)

}
