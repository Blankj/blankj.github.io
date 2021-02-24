package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}

	select {
	case <-ch: // 接收任意数据
		...
	case d := <-ch: // 接收变量
	case ch <- 100: // 发送数据
	default:
		没有操作情况
	}

}
