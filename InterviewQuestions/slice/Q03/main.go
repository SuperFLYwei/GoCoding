package main

import "fmt"

/*
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}


 */

/*
copy()
copies min(len(dst), len(src)) elements
 */
func copy_test() {
	src := []string{"a", "b", "c"}
	dst := make([]string, len(src))

	copy(dst, src)

	fmt.Printf("source slice: %[1]v, address: %[1]p\n", src)
	fmt.Printf("source slice: %[1]v, address: %[1]p\n", dst)
}

func copy_test2() {
	src := []string{"a", "b", "c"}
	var dst []string

	dst = append(dst, src...)

	fmt.Printf("source slice: %[1]v, address: %[1]p\n", src)
	fmt.Printf("source slice: %[1]v, address: %[1]p\n", dst)
}

// shallow copy
func shallow_copy() {
	src := []string{"a", "b", "c"}
	dst := src

	fmt.Printf("source slice: %[1]v, address: %[1]p\n", src)
	fmt.Printf("source slice: %[1]v, address: %[1]p\n", dst)
}


func main() {
	copy_test()
	fmt.Println()
	shallow_copy()
}
