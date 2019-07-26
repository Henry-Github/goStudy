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

	d := [5]int{1, 2, 3, 4, 5}
	var e [5]int
	e[0] = d[0]

	fmt.Println(d)
	fmt.Println(e)

}
