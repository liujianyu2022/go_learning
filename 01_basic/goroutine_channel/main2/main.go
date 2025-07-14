package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	// 主线程
	for i := 0; i < 3; i++ {
		fmt.Println("main = ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	// 由于存在主死从随的特性。即主线程的程序一旦执行完毕，协程会立马停止。
	// 为了让协程代码能够顺利执行，采用waitGroup等待
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)										// 每开启一个协程，计数器增加1
		go func (num int){
			defer waitGroup.Done()								// 每结束一个协程，计数器减少1	   defer 关键字后的语句会被压入一个栈，在函数执行完之后再执行
			fmt.Println("num = ", strconv.Itoa(num))
		}(i)
	}

	waitGroup.Wait()											// 一直等待，直到计数器变为0
}
