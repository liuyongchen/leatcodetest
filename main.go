package main

import (
	"fmt"
)

type Persion struct {
	Name string
	Age  int64
}

type Student struct {
	Persion Persion
	Grade   string
}

func main() {
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
	a := 10
	happy := isHappy(a)
	fmt.Println(happy)
	v := '('
	fmt.Printf("%T", v)
}
