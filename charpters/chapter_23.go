package main

import "fmt"

// channels

func inChan(ch chan string, name string) {
	ch <- name
}

func outChan(ch chan string) string {
	return <-ch
}

func main() {
	ch := make(chan string, 2)
	go func(c chan string) {
		c <- "test"
	}(ch)

	for {
		name := <-ch
		fmt.Println(name)
	}

}
