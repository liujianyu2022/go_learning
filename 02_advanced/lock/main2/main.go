package main

import "sync"

type Person struct {
	mutex  sync.Mutex
	salary uint
	level  uint
}

// 修改的时候加上互斥锁，保证修改的时候只有一个协程操作数据
// 写锁是互斥锁，一旦加上了写锁，其他协程就不能加写锁或者读锁了
func (person *Person) promote() {
	person.mutex.Lock()
	defer person.mutex.Unlock()

	person.salary += 1000
	person.level += 1
}

// 并发的读取数据，但是写操作需要被阻止
// 读锁是共享锁，当加上了读锁之后，其他协程可以继续加读锁，但是不能加写锁
func (person *Person) print(){

}

func main() {

}
