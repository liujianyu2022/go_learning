package main

import (
	"fmt"
	"sync"
)

type Person struct {
	mutex  sync.Mutex // 互斥锁，只有Lock和Unlock方法，当读写操作都需要互斥时使用
	name  string
	salary uint
	level  uint
}

// 修改的时候加上互斥锁，保证修改的时候只有一个协程操作数据，写锁是互斥锁，一旦加上了写锁，其他协程就不能加写锁或者读锁了
func (person *Person) promote(waitGroup *sync.WaitGroup) {
	person.mutex.Lock()
	defer person.mutex.Unlock()

	person.salary += 1000
	person.level += 1

	waitGroup.Done()
}

// 由于 sync.Mutex 只有 Lock 和 Unlock 方法，所以读锁和写锁是一样的效果，即加了读锁之后，其他协程也不能加读锁或者写锁
func (person *Person) print(waitGroup *sync.WaitGroup) {
	person.mutex.Lock()
	defer person.mutex.Unlock()

	fmt.Println("salary = ", person.salary, ", level = ", person.level)

	waitGroup.Done()
}

func main() {
	person := Person{
		salary: 100000,
		level: 10,
	}

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(6)
	go person.promote(&waitGroup)
	go person.promote(&waitGroup)
	go person.promote(&waitGroup)

	go person.print(&waitGroup)
	go person.print(&waitGroup)
	go person.print(&waitGroup)

	waitGroup.Wait()
}