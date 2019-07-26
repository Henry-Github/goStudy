package main

import (
	"fmt"
)

// slices
func main() {
	s := make([]int, 5)
	fmt.Println(len(s))
	for i := 0; i < 5; i++ {
		s[i] = i
	}
	ss := make([]int, 5)
	fmt.Println(len(s))
	for i := 0; i < 5; i++ {
		ss[i] = i + 5
	}
	var sss []int
	fmt.Println(len(sss))
	for i := 0; i < 10; i++ {
		sss = append(sss, i)
	}
	fmt.Println(len(sss))

	fmt.Println(s)
	fmt.Println(ss)
	fmt.Println(sss)

}
