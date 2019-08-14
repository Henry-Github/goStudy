package main

import "fmt"

// channels

func myClourse() func() int {
	var x = 0
	return func() int {
		x = x + 1
		return x
	}
}

func main() {
	fmt.Println(myClourse())
	fmt.Println(myClourse())
	fmt.Println(myClourse())
	fmt.Println(myClourse())
	fmt.Println(myClourse())
	fmt.Println(myClourse())
	fmt.Println(myClourse())

}
