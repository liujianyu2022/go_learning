package main

import "time"

// 在 Go 语言中，time 包下的 Timer 类型用于 在指定时间之后执行某个动作。

// 基础用法，当作定时器使用，类似于 setTimeout
func test1(duration time.Duration) {

	timer := time.NewTimer(duration)

	// 阻塞等待，直到定时器到期
	// timer.C 是一个 只读的 channel，它的作用是 在定时器到期时接收一个信号（时间值）
	<-timer.C

	println("test1 定时器到期，执行相应操作")
}

// 超时控制（常见于网络请求/并发任务）
func test2(duration time.Duration) {
	timer := time.NewTimer(duration)

	// Stop() 用于停止还没触发的定时器。
	// 否则即使程序退出，Timer 的底层 goroutine 依然存在，会造成轻微资源泄漏。
	defer timer.Stop() 		

	done := make(chan bool)

	// 并发逻辑：开启一个 goroutine，模拟一个耗时操作
	go func() {
		time.Sleep(1 * time.Second)
		done <- true
	}()

	// 如果任务先完成（1 秒 < 2 秒定时器），执行 done 分支；
	// 如果任务没完成且超过 2 秒，执行 timer.C 分支（超时）。
	select {
	case <-done:
		println("test2 耗时操作完成，定时器被取消")
	case <-timer.C:
		println("test2 定时器到期，执行相应操作")
	}
}

func main() {
	go test1(3 * time.Second)
	go test2(2 * time.Second)

	time.Sleep(6 * time.Second)
}
