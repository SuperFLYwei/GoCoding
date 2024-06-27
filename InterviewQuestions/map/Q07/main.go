package main

import (
"fmt"
"reflect"
)

func main() {
	// 创建两个相同的map对象
	map1 := map[string]int{
		"apple":  1,
		"orange": 2,
		"banana": 3,
	}
	map2 := map[string]int{
		"banana": 4,
		"orange": 2,
		"apple":  1,
	}

	// 使用reflect.DeepEqual来比较两个map是否相等
	result := reflect.DeepEqual(map1, map2)

	// 输出结果
	if result {
		fmt.Println("两个map对象是相等的")
	} else {
		fmt.Println("两个map对象是不相等的")
	}
}
