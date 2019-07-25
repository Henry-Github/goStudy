package main

import (
	"fmt"
)

// loop

func main() {
	var i int32 = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	for j := 20; j < 30; j++ {
		fmt.Println(j)
	}

}
