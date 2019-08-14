package main

import (
	"fmt"
	"time"
)

// timeouts 超时
// 超时检测是一个非常重要的方法
// 在做api请求时 如果api是带有返回值的 一般情况下需要做返回结果检查
// 返回结果ok符合正常请求继续处理 否则根据异常情况做相应的处理 这里使用panic 异常退出

func main() {

	c := make(chan int)

	go func(c chan int) {
		time.Sleep(time.Second * 3)
		c <- 1
		close(c)
	}(c)

	for {
		select {
		case res, ok := <-c:
			fmt.Println(ok)
			if ok {
				fmt.Println("get data: ", res)
			} else {
				fmt.Println("channel closed, can not get data")
			}

		case <-time.After(time.Second * 2):
			panic("timeout")

		}
	}

}
