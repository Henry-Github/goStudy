package main

import (
	"fmt"
	"time"
)

// 闭包、匿名函数
// 闭包：一个拥有许多变量和绑定了这些变量的表达式(通常为函数)
// 闭包使用全局变量时候， 新的表达式对象，调用的时候，将会
// 1. 包含一个函数
// 2. 该函数调用原函数变量
// 3. 返回该函数

func executeFunc() func() int {
	i := 0
	return func() int {
		i++
		//fmt.Printf("this is %d\n ", i)
		return i
	}
}

func executeFunc1() func() int {
	return func() int {
		i := 0
		i++
		return i
	}
}

var p = 0

func executeFunc2() func() int {
	return func() int {
		p++
		return p
	}
}

func main() {

	go func() {
		exe := executeFunc2()
		fmt.Println(exe(), 1)
		fmt.Println(exe(), 2)
		fmt.Println(exe(), 3)
		fmt.Println(exe(), 4)
	}()
	go func() {
		newExe := executeFunc2()
		fmt.Println(newExe(), 5)
		fmt.Println(newExe(), 6)
		fmt.Println(newExe(), 7)
		fmt.Println(newExe(), 8)
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1000)
	}
}
