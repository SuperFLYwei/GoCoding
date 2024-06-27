package main

import (
	"fmt"
	"sync"
)

func main() {
	var capId sync.Map

	capId.Store("Beijing", 88)
	capId.Store("London", 80)
	capId.Store("Tokyo", "89") // 注意value是字符串

	fmt.Println(capId.Load("Beijing"))
	fmt.Println(capId.Load("Tokyo"))

	fmt.Println()

	value, _ := capId.Load("Beijing")
	fmt.Println(value)

	fmt.Println()

	capId.Delete("Tokyo")
	capId.Range(func(k, v interface{}) bool {
		fmt.Printf("%s对应的id是%d\n", k, v)
		return true
	})
}

// output:

// 88 true
// 89 true
// <nil> false
//
// 88
//
// Beijing对应的id是88
// London对应的id是80

