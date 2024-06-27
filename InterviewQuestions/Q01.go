package main

import "fmt"

func Q01() {
	arr := [5]int{1,2,3,4,5}
	m := make(map[int]int, 5)
	for i , v := range arr {
		m[i] = v
	}
	fmt.Println(m)
}

func main() {
	Q01()
}
