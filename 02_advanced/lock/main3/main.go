package main

import (
	"fmt"
	"sync"
)

type Person struct {
	mutex sync.RWMutex
	salary uint
	level uint
}

func (person *Person) promote(waitGroup *sync.WaitGroup) {
	person.mutex.Lock()
	defer person.mutex.Unlock()

	person.salary += 1000
	person.level += 1

	waitGroup.Done()
}

func (person *Person) print(waitGroup *sync.WaitGroup) { 
	person.mutex.RLock()
	defer person.mutex.RUnlock()

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
