package main

import (
	"fmt"
	"time"
)

// switch

func checkType(i interface{}) {
	switch i.(type) {
	case int8:
	case int32:
	case int16:
	case int64:
		fmt.Println("type of int")
	case string:
		fmt.Println("type of string")
	case bool:
		fmt.Println("type of bool")
	}
}

func main() {
	var name = "fds"
	switch name {
	case "henry":
		fmt.Println(" this is henry")
	case "john":
		fmt.Println("this is john")
	default:
		fmt.Println("this is unknown")
	}

	var day = time.Now().Weekday()
	switch day {
	case time.Wednesday:
		fmt.Println("today is Wednesday")
	default:
		fmt.Println(day)
	}

	var now = time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("morning")
	default:
		fmt.Println("afternoon")
	}

	var b bool = true
	var s string = "rewr"
	var a int64 = 1
	checkType(b)
	checkType(s)
	checkType(a)

}
