package main

import (
	"fmt"
	"unsafe"
)

// 空结构体，不会占用任何的内存空间
// 空结构体的地址均相同（当不被包含在其他结构体中的时候）
// 所有 struct{} 实例共享同一内存地址（如 zerobase），避免重复分配。
type K struct{}

// 慎用这种写法
// 这样空结构体会占用内存空间了
type F struct {
	member K
}

func main() {

	var k1 = K{}
	var k2 = K{}
	fmt.Printf("address of k1 = %p\n", &k1)         // 有地址
	fmt.Printf("address of k2 = %p\n", &k2)         // 有地址		并且 k1 和 k2 的地址相同，在Go中，所有长度为0的空结构体指向的地址都是相同的
	fmt.Println("size of k1 = ", unsafe.Sizeof(k1)) // 长度为0		就好比有门牌号，但是房子面积为0

	// 空结构体作用1
	// map + 空结构体 = set	   map[T]struct{}
	// map: key-value    set: key-nil  也就是说set只需要key，不需要value
	set1 := map[string]struct{}{}
	set1["a"] = struct{}{}

	set2 := make(map[string]struct{})
	set2["b"] = struct{}{}

	if _, existed := set1["a"]; existed {
		fmt.Println("a exists in set1")
	}

	// 空结构体作用2
	// 实现一个“信号”通道，仅用于通知而非传递数据。
	ch := make(chan struct{})
	ch <- struct{}{}				// 发送空结构体，表示发出了一个信号
	<- ch							// 接收信号，关心具体的值
}
