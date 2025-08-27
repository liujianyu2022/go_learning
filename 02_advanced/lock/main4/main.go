package main

import (
	"fmt"
	"sync"
)

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
			once.Do(func() {test(&num)})
		}()
	}

	waitGroup.Wait()

	fmt.Println("num = ", num)
}
