package main

import (
	"fmt"
	"time"
)

// 无缓冲的管道，必须要有 同时存在 的发送方和接收方，才能完成一次传递。不然就会阻塞
func demo1(){
	var channel1 chan string = make(chan string)    
	var channel2 chan string = make(chan string, 0)

	// 由于没有接收方，下面两行会阻塞程序执行
	channel1 <- "a"
	channel2 <- "b"

	fmt.Println("demo1：接收方不是同时存在的，", <- channel1, <- channel2)

	fmt.Println("demo1：无缓冲的管道，会阻塞程序执行，这句话无法输出")
}

// 无缓冲通道：同时存在发送方和接收方，因此不会阻塞
func demo2(){
	channel := make(chan string)

	// 启动一个协程作为接收方
	// 由于同时存在发送方和接收方，因此不会阻塞
	go func(){
		fmt.Println("demo2：有消费者等待，", <- channel)
	}()

	channel <- "demo2"
}

// 有缓冲的管道，发送方在缓冲区没满时不会阻塞，接收方在缓冲区没空时不会阻塞。
func demo3(){
	channel := make(chan int, 2) 		

	channel <- 0
	channel <- 1			// 前面两个进入缓冲区，程序不阻塞

	fmt.Println("demo3：有缓冲的管道，在缓冲区满之前是不会阻塞的")

	// 缓冲区满了，下面这行会阻塞程序执行
	channel <- 2
	fmt.Println("demo3：有缓冲的管道，缓冲区满之后会阻塞，因此这里无法输出")
}

func demo4(){
	channel := make(chan int, 2)

	channel <- 0
	channel <- 1			// 前面两个进入缓冲区，程序不阻塞

	fmt.Println("demo4：有缓冲的管道，在缓冲区满之前是不会阻塞的")

	go func(){
		fmt.Println("demo4：有消费者等待，", <- channel)
	}()

	channel <- 2
}

func main() {
	go demo1()
	go demo2()
	go demo3()
	go demo4()

	time.Sleep(5 * time.Second)
}

// 总结来说：
// 无缓冲 channel 必须成对收发，否则阻塞。
// 有缓冲 channel 不能超容量，否则阻塞
// Go 的 channel 设计哲学就是“零丢失”——只要写入成功，就保证消息一定会被某个接收方读出来。不会像消息队列里那样默认有丢弃策略。