package main

import (
	"fmt"
	"sync"
)

// make 和 new
/*
1. make 只用于 channel slice map 创建
	func make(t Type size..IntegerType) Type
   make 返回的是类型本身， 使用时为引用使用
2. new 参数只有一个 返回值为指针 且初始化为零值
	func new(Type) *Type
*/

type Person struct {
	lock sync.Mutex
	name string
	age  int
}

func main() {
	a := new(int) //初始化 且赋值为0
	fmt.Println(*a)
	*a = 10
	fmt.Println(*a)
	b := make([]string, 2, 10)
	fmt.Println("b:", b)

	p := new(Person)
	fmt.Println(p.name)
	fmt.Println(p.age)

	var pp Person
	fmt.Println(pp.name)
	fmt.Println(pp.age)
}
