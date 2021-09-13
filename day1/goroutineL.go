package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 通道 channel 可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

func main() {

	//  go 语句开启一个新的运行期线程 同一个程序中的所有 goroutine 共享同一个地址空间
	go say("world")
	say("hello")
}
