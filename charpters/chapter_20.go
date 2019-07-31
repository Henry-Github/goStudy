package main

import (
	"fmt"
)

// goroutine

func testGo(test string) {
	for i := 0; i < 4; i++ {
		fmt.Printf("this comes from %s, loop: %d\n", test, i)
	}
}

func testGoroutineParm(i int) {
	fmt.Println("param: ", i)
}

func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("param: ", i)
		}(i)
		//go testGoroutineParm(i)
	}
	//go func() {}()
	//
	//
	//testGo("test")
	//go testGo("goroutine")
	//go func(going string) {
	//	for i:=0;i<4; i++ {
	//		fmt.Printf("this comes from %s, loop: %d\n ", going, i)
	//	}
	//
	//}("going")
	fmt.Scanln()
	fmt.Println("done")

}
