package main

import "fmt"

// variadic functions

func add1(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func add2(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func main() {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	fmt.Println(add1(s...))
	fmt.Println(add2(s))
}
