package main

import "fmt"

func Q01()  {
	m := make(map[int]int)
	m[1] = 2
	m[1] = 3
	for _, v := range m {
		fmt.Println(v)
	}
}

func main() {
	Q01()
}
