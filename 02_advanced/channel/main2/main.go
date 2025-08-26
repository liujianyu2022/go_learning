package main

import (
	"fmt"
	"time"
)

// 通过共享内存通信
// 这种方式就是典型的 共享内存 + 轮询，缺点是：
// CPU 空转（一直在 for 循环里检查条件，很浪费）。
// 存在数据竞争问题（多个 goroutine 同时读写 num 没有加锁）
func demo1(pointer *int) {
	for {
		if *pointer == 1 {
			fmt.Println("demo1：通过共享内存的方式通信")
			break
		}
	}
}

// 通过 channel 通信
// 天然同步，不会有数据竞争，不浪费 CPU
func demo2(channel chan int){
	if <-channel == 1 {
		fmt.Println("demo2：通过管道的方式通信")
	}
}

func main() {
	var num int = 0
	var channel chan int = make(chan int, 1)

	go demo1(&num)
	time.Sleep(3 * time.Second)
	num = 1

	go demo2(channel)
	time.Sleep(3 * time.Second)
	channel <- 1

	time.Sleep(3 * time.Second)
}
