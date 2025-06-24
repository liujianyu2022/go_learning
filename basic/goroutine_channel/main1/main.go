package main

import "fmt"

/*
3. 管道本质是一个队列，数据先进先出
4. 管道自身就是线程安全的
*/

func DefineChannel(){
	// 管道的定义   var 变量名 chan 数据类型        		
	var stringChannel chan string = make(chan string, 4)		// 数据类型	管道容量   可读可写

	// var stringChan chan<- string = make(chan string, 4)		// 数据类型	管道容量   只可以写
	// var stringChan <-chan string = make(chan string, 4)		// 数据类型	管道容量   只可以读

	fmt.Println("stringChan = ", stringChannel)				    // 管道属于引用数据类型  打印出来的是地址值

	stringChannel <- "hello"
	stringChannel <- "world"
	stringChannel <- "!"

	fmt.Printf("管道的实际长度 = %d，管道的容量 = %d \n", len(stringChannel), cap(stringChannel))		// 3, 4

	str1 := <- stringChannel
	str2 := <- stringChannel
	str3 := <- stringChannel

	fmt.Printf("str1 = %s, str2 = %s, str3 = %s \n", str1, str2, str3)
}

// 管道关闭后，不能再向管道中存入数据，但可以从管道中读取数据
func CloseChannel(){
	stringChannel := make(chan string, 4)

	stringChannel <- "hello"
	stringChannel <- "world"

	close(stringChannel)

	// stringChannel <- "!"              // 管道关闭后再写输入就会报错

	// 从管道中读数据
	str1 := <- stringChannel
	str2 := <- stringChannel

	fmt.Printf("str1 = %s, str2 = %s \n", str1, str2)	
}

// 管道的遍历，只能使用for-range  并且在遍历管道之前，需要关闭管道，否则会报错 deadlock
func ForRange(){
	stringChannel := make(chan string, 5)

	stringChannel <- "hello"
	stringChannel <- "world"

	close(stringChannel)							// 遍历管道之前，需要关闭管道，否则会报错 deadlock

	for value := range stringChannel {				// for-range遍历管道的时候，只有 value，没有key
		fmt.Printf("value = %s \t", value)
	}
}


func main() {
	fmt.Println(" ------------ 管道的定义 --------- ")
	DefineChannel()
	
	fmt.Println(" ------------ 管道的close --------- ")
	CloseChannel()

	fmt.Println(" ------------ 管道的遍历 --------- ")
	ForRange()
}
