package main

import "fmt"

type MyInterface interface {
	Print()
}

type MyStruct struct{}
func (ms MyStruct) Print() {}

func main() {
	x := 1
	var y interface{} = x
	var s MyStruct
	var t MyInterface = s
	fmt.Println(y, t)
}