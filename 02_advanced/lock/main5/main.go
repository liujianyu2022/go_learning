package main

import "sync"

// 使用 sync.Once 实现单例模式

type Singleton struct {
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {

	// once.Do 确保下面的函数只会被执行一次
	once.Do(func() {
		instance = &Singleton{}
	})
	
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 == s2 {
		println("s1 和 s2 是同一个实例")
	}
}
