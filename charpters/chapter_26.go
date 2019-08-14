package main

import (
	"fmt"
	"time"
)

// channel select

func receiveMultiChan(c1 <-chan string, c2 <-chan string) {
	for {
		time.Sleep(time.Second)
		select {
		case rec := <-c1:
			fmt.Printf("receiving from c1 content is : %s\n", rec)
		case rec := <-c2:
			fmt.Printf("receiving from c2 content is : %s\n", rec)

		}
	}
}

func waitInput() {
	var input string
	fmt.Scanln(&input)
	fmt.Println("scan enter:", input)
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)
	go func(c chan string, name string) {
		for i := 0; ; i++ {
			c <- name
		}
	}(c1, "c1")
	go func(c chan string, name string) {
		for i := 0; ; i++ {
			c <- name
		}
	}(c2, "c2")
	go receiveMultiChan(c1, c2)
	waitInput()

}
