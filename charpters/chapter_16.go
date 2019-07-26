package main

import "fmt"

// 指针 pointer
// & 取到变量的地址 *& = new val  改变实际值
// 常规函数只是引用 不会改变原始值

func zeroval(val int) {
	fmt.Println("值复制：", &val)
	val = 0
}

func zeroptr(ptr *int) {
	fmt.Println("指针引用制：", ptr)
	*ptr = 0
}

func arrayVal(val []int) {
	val[0] = 10
}

func arrayPtr(val *[]int) {
	a := val
	fmt.Println(a)
}

func main() {
	i := 10
	fmt.Println(i)
	fmt.Println("原始值地址：", &i)
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)

	var a []int
	for i := 0; i < 4; i++ {
		a = append(a, i)
	}
	fmt.Println("数组原始值地址:", &a)
	arrayVal(a)
	fmt.Println(a)
	arrayPtr(&a)
}
