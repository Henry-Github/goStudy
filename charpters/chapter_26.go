package main

import (
	"fmt"
	"time"
)

// channel select

type Id struct {
	id   int
	name string
}

func receiveMultiChan(c1 <-chan Id, c2 <-chan Id) {
	for {
		time.Sleep(time.Second)
		select {
		case rec := <-c1:
			fmt.Printf("receiving from c1 content is : %s-%d\n", rec.name, rec.id)
		case rec := <-c2:
			fmt.Printf("receiving from c2 content is : %s-%d\n", rec.name, rec.id)
		}
	}
}

func waitInput() {
	var input string
	fmt.Scanln(&input)
	fmt.Println("scan enter:", input)
}

func main() {
	c1 := make(chan Id, 1)
	c2 := make(chan Id, 1)
	go func(c chan Id, name string) {
		for i := 0; ; i++ {
			c <- Id{id: i, name: name}
		}
	}(c1, "c1")
	go func(c chan Id, name string) {
		for i := 0; ; i++ {
			c <- Id{id: i, name: name}
		}
	}(c2, "c2")
	go receiveMultiChan(c1, c2)
	waitInput()

}
