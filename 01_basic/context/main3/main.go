package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // 确保资源被释放

	go func() {
		for {
			select {
			case <-ctx.Done(): // 当到达截止时间之后， ctx.Done()会接收到取消信号，退出协程
				println("收到取消信号，退出协程, error = ", ctx.Err())
				return
			default:
				if deadline, ok := ctx.Deadline(); ok {
					remaining := time.Until(deadline)
					fmt.Printf("剩余时间: %.1f 秒\n", remaining.Seconds())
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	<-ctx.Done() // 主协程等待，直到到达截止时间

	time.Sleep(1 * time.Second) // 确保子协程有时间打印退出信息

	println("main 退出")
}
