package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建一个map
	m := map[string]int{
		"apple":  1,
		"orange": 2,
		"banana": 3,
	}

	// 将map的键存储在一个切片中
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 按照排序后的键的顺序读取map的值
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
