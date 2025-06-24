package main

import (
	"fmt"
	"sync"
	"time"
)

// chan 是一种通道类型，用于在不同的 Goroutine（协程）之间安全地传递数据，是实现并发编程中同步和通信的核心机制
// chan 是 Go 的内置类型，属于引用类型（底层通过指针实现）,遵循先进先出（FIFO）的规则，确保数据按发送顺序接收
// 阻塞机制：发送数据到通道时，如果通道已满（有缓冲且满），发送方会阻塞，直到数据被接收。从通道接收数据时，如果通道为空，接收方会阻塞，直到有数据可读
// 方向性：通道可以声明为单向（只读 <-chan 或只写 chan<-），通常用于函数参数限制权限

// 同步通道（无缓冲区通道）。 发送和接收操作必须同时就绪，否则会阻塞 Goroutine
// 确保 Goroutine 同步：强制发送和接收双方同步执行，常用于等待任务完成或信号通知
func RunChannelDemo1() {
	ch := make(chan string) 							// 同步通道, 无缓冲区，通道的容量为 0。发送方将数据直接交给接收方，没有中间存储

	var waitGroup sync.WaitGroup
    waitGroup.Add(1)									// 等待 1 个 Goroutine 完成
	
	go func() {
		defer waitGroup.Done()							// 确保 Goroutine 结束时通知 WaitGroup。    WaitGroup 必须和 Done() 配对使用，否则会死锁。

		fmt.Println("Goroutine 发送数据...")
		ch <- "Hello" 									// 发送数据，阻塞直到主 Goroutine 接收。     必须等待另一个 Goroutine 执行接收操作（<-ch），否则发送方阻塞。    
		fmt.Println("发送完成")
	}()

	time.Sleep(2 * time.Second) 						// 模拟耗时操作
	msg := <-ch                 						// 接收数据，解除阻塞。                      必须等待另一个 Goroutine 执行发送操作，否则接收方阻塞
	fmt.Println("主程序收到:", msg)

	waitGroup.Wait() // 等待发送方 Goroutine 完成
}

// 异步通道（有缓冲通道）
// 异步非阻塞：发送操作在缓冲区未满时立即完成，否则阻塞。接收操作在缓冲区非空时立即完成，否则阻塞。允许短暂的数据堆积。

func RunChannelDemo2() {
	ch := make(chan int, 2) 							// 异步通道，有缓冲区：通道容量 > 0，这里的缓冲容量为 2

	// 生产者（快速发送）
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i										// 若缓冲区未满，数据存入缓冲区，发送方继续执行；若缓冲区已满，发送方阻塞。
			fmt.Printf("发送 %d \n", i)
		}
		close(ch)										// 关闭channel。只有当确实没有数据发送，或者想显式的结束range循环之类的，才去关闭channel。
	}()

	// 消费者（慢速接收）
	for value1 := range ch {							// 隐式的接收操作，它会不断从通道 ch 中读取数据，直到通道被关闭 close(ch)
		time.Sleep(1 * time.Second) 					// 模拟耗时操作
		fmt.Printf("Received: %d\n", value1)
	}

	// 下面是等价操作，显示的接收操作
	// for {
	// 	value2, ok := <-ch  							// 接收数据，并检查通道是否关闭
	// 	if !ok {       									// 如果通道已关闭，退出循环
	// 		break
	// 	}
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Printf("Received: %d\n", value2)
	// }
}

