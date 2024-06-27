package main

import "fmt"

// Output:
//[1 2 3]
//[1 2 3 5]
//[1 2 3 5]
func AppendDemo() {
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}



func AppendDemo2() {
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	x = append(x, 4)
	z := append(x, 5)
	fmt.Println(x)
	fmt.Println(x)
	fmt.Println(z)
}


func main() {
	AppendDemo()
	AppendDemo2()
}

