package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func add1(pointer *int64) {
	*pointer += 1
}

func add2(pointer *int64){
	atomic.AddInt64(pointer, 1)
}

func main() {
	var num1 int64 = 0
	var num2 int64 = 0

	for i := 0; i < 1000; i++ {
		go add1(&num1)
		go add2(&num2)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)
}
