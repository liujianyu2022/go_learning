package main

import (
	"fmt"
	"time"
)

// 向 channel 中写入
func test1(channel *chan int){
	*channel <- 1
}

// 从 channel 中读取
func test2(channel *chan int){
	result := <- *channel
	fmt.Println("result = ", result)
}

func main() {
	channel1 := make(chan int, 5)
	channel2 := make(chan int)

	go test1(&channel1)				// 向 channel1 中写入数据
	go test2(&channel2)				// 从 channel2 中读取数据

	// select 结合 default 分支，可以实现非阻塞的读写
	// 注意：当有多个 case 同时可以执行时，Go runtime 会随机选择一个执行

	// 即如果没有 channel 可操作就执行 default，而不会阻塞当前 goroutine。
	select {
	// 如果channel1成功读取到数据，就会进入这个case
	// 此时 channel1 为空缓冲，无法读取，如果没有select，那么就会阻塞等待
	case <- channel1:
		fmt.Println("in channel1")

	// 如果往channel2中成功写入数据，就会进入这个case
	// 此时 channel2 是无缓冲，没有接收者，无法写入。如果没有 select，那么就会阻塞等待
	case channel2 <- 1:
		fmt.Println("in channel2")

	default:
		fmt.Println("in default")
	}

	fmt.Println("in main")

	time.Sleep(2 * time.Second)
}
