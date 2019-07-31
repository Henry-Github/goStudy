package main

import (
	"errors"
	"fmt"
)

// go error

func checkError(i int) (res int, err error) {
	switch {
	case i > 10:
		return 10, errors.New("param over 10")
	case i < 0:
		return 0, errors.New("param less 0")
	default:
		return i, nil

	}
}

func main() {
	var test = [6]int{-1, -4, 1, 2, 4}
	for _, v := range test {
		if _, ok := checkError(v); ok != nil {
			fmt.Println("error element: ", v)
		}
	}
}
