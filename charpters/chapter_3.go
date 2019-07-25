package main

import "fmt"

func main() {
	var a = 10
	fmt.Println(a)
	var b = a
	b = 9

	fmt.Println(a)
	fmt.Println(b)
}
