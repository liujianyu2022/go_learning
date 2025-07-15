package main

import (
	"fmt"
	"unsafe"
)

// 在 Go 语言中，基本数据类型的长度（占用的字节数）取决于具体的平台（主要是 32 位或 64 位系统）
// 在计算机中，1 字节（Byte）等于 8 比特（bit）    1 byte = 8 bit

// 类型		32 位系统		64 位系统
// uint8	1 字节			1 字节
// uint16	2 字节			2 字节
// uint32	4 字节			4 字节
// uint64	8 字节			8 字节
// int		4 字节			8 字节
// uint		4 字节			8 字节
// uintptr	4 字节			8 字节

// float32	4 字节			4 字节
// float64	8 字节			8 字节

// bool		1 字节			1 字节
// byte		1 字节			1 字节

// string	8 字节			16 字节			（32位： 4字节指针 + 4字节长度字段； 64位： 8字节指针 + 8字节长度字段）

// 指针		4 字节			8 字节

func main(){
	var num1 uint8 = 10
	var num2 uint16 = 10
	var num3 uint32 = 10
	var num4 uint64 = 10
	
	var num5 uint = 10

	fmt.Println("size of unit8 = ", unsafe.Sizeof(num1))				// 返回变量的字节数，1
	fmt.Println("size of unit16 = ", unsafe.Sizeof(num2))				// 2
	fmt.Println("size of unit32 = ", unsafe.Sizeof(num3))				// 4
	fmt.Println("size of unit64 = ", unsafe.Sizeof(num4))				// 8
	fmt.Println("size of unit = ", unsafe.Sizeof(num5))					// 和系统相关，在64位系统中，uint默认是 uint64

	var ptr1 *uint8 = &num1
	var ptr2 *uint16 = &num2
	fmt.Println("size of *uint8 = ", unsafe.Sizeof(ptr1))				// 指针长度和系统相关，32位系统中位4字节， 64位系统中为8字节
	fmt.Println("size of *uint16 = ", unsafe.Sizeof(ptr2))

	var flag bool = true
	fmt.Println("size of bool = ", unsafe.Sizeof(flag))					// 1

	var str string = "hello"
	fmt.Println("size of str = ", unsafe.Sizeof(str))					// 16
}
