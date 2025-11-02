package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 确保资源被释放

	go func() {
		for {
			select {
			case <-ctx.Done(): // 当超时之后， ctx.Done()会接收到取消信号，退出协程
				println("收到取消信号，退出协程, error = ", ctx.Err())
				return
			default:
				println("协程正在运行中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	<-ctx.Done() // 主协程等待，直到超时

	time.Sleep(1 * time.Second)			// 确保子协程有时间打印退出信息

	println("main 退出")
}
