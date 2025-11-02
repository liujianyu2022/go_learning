package main

import (
	"fmt"
	"sync"
)

// sync.Once 保证某段代码只执行一次

func test(pointer *int) {
	*pointer += 1
}

func main() {
	num := 0
	once := sync.Once{}
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(3)

	for i := 0; i < 3; i++ {
		go func(){
			defer waitGroup.Done()
			once.Do(func() {test(&num)})			// 保证 test 函数只会被执行一次
		}()
	}

	waitGroup.Wait()

	fmt.Println("num = ", num)
}
