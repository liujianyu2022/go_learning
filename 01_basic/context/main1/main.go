package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(){
		for {
			select {
			case <- ctx.Done():		// 当 cancel被调用的时候， ctx.Done()会接收到取消信号，退出协程
				fmt.Println("收到取消信号，退出协程, error = ", ctx.Err())
				return
			default: 
				fmt.Println("协程正在运行中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("准备取消协程...")
	cancel()								// 发送取消信号

	time.Sleep(1 * time.Second)
	fmt.Println("main 退出")
}