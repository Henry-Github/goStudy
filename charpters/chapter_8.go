package main

import "fmt"

// array

func main() {
	var a [5]int
	a[4] = 10
	fmt.Println(a)
	var b = 10
	var c = &b

	fmt.Println(a)
	fmt.Println(&b)
	fmt.Println(*c)

	p := &a
	fmt.Println(p[4])

}
