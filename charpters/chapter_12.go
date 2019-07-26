package main

import "fmt"

// function

func myAdd(x int, y int) int {
	return x + y
}

func void() (int, int) {
	return 2, 4
}

func main() {
	c := myAdd(1, 2)
	fmt.Println(c)
	fmt.Println(void())
}
