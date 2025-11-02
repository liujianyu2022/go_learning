package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 启动100个协程，分别计算两个随机整数之和，并通过管道通知主协程结果已完成；
// 主协程通过管道接收通知，统计完成的加法运算数量，直到全部完成后退出程序；

// 计算两个整数之和，并通过管道通知结果已完成；
// 使用空结构体作为信号，节省内存，因为空结构体不占用任何内存空间；
// chan<- 这是一个只可以写入的管道
func add(x, y int, channel chan<- struct{}) {
	fmt.Printf("%d + %d = %d\n", x, y, x + y)
	channel <- struct{}{}
}

func main() {
	channel := make(chan struct{}, 100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			// 为每个 goroutine 创建独立的随机数生成器
			localRand := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
			
			x := localRand.Intn(100)
			y := localRand.Intn(100)
			add(x, y, channel)
		}(i)
	}

	count := 0
	for {
		select {
		case <- channel:
			count++
			if count == 100 {
				fmt.Println("所有加法运算完成")
				return
			}
		}
	}
}
