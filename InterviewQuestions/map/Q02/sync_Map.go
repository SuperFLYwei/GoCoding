package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Store(i+1, i)
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			v, _ := m.Load(i)
			fmt.Printf("使用map的key：%d,取到map的value:%d\n", i, v)
		}(i + 1)
	}
	time.Sleep(time.Second)
}

