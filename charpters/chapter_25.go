package main

import (
	"fmt"
	"time"
)

// channel direction
// 指定channel为单方向 或者双向

func sender(sender chan<- string, content string) {
	for i := 0; ; i++ {
		sender <- content
	}
}

func receiver(receiver <-chan string, name string) {
	for {
		rec := <-receiver
		time.Sleep(time.Second)
		fmt.Printf("receive_%s content: %s\n", name, rec)
	}
}

func waitKey() {
	var input string
	fmt.Scanln(&input)
	fmt.Println("entered: ", input)

}

func main() {
	c := make(chan string)
	go sender(c, "sender-1")
	go sender(c, "sender-2")
	go receiver(c, "1")
	go receiver(c, "2")
	go receiver(c, "3")
	waitKey()

}
