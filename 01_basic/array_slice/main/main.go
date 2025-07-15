package main

import (
	"fmt"
	"go_learning/01_basic/array_slice/tools"
)

// 数组定义		var 数组名 [数组长度]数据类型		var array [5]int	 长度也是数据类型的一部分   [3]int  [5]int 数据类型不同
// 切片：	  	var 切片名 []数据类型			   var slice []int 		不用指定长度，对数组一个连续片段的引用

// 切片是一个结构体，包括三个属性 {底层数组的指针，切片长度，切片长度}

/*
   数组：长度固定不可变，并且长度也是数组数据类型的一部分，比较呆板								golang中，数组属于基本数据类型
   切片：对数组中一个连续片段的引用，因此切片属于引用数据类型，对切片的改变会影响底层数组的值
*/

func main() {
	fmt.Println("------------ 下面展示一维数组 ------------")

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr2 = [5]int{1, 2, 3, 4, 5}
	var arr3 = [...]int{1, 2, 3, 4, 5}											// 自动推断数组长度
	var arr4 = [...]string{"hello", "你好", "世界"} 							// 注意：对于数组来说，len函数返回的是数组的元素个数，因此遍历数组的时候汉字不会出现乱码。

	// 注意：由于长度也是数组数据类型的一部分，因此 [5]int 和 [3]int 是不同的数据类型 
	tools.ForRange1(arr1)
	tools.For1(arr2)
	tools.For1(arr3)

	tools.For2(arr4)
	tools.ForRange2(arr4)

	// 编写遍历函数的时候， [5]int 和 [3]int 需要分别编写，十分不方便
	// 因此遍历函数使用切片作为形参类型即可
	tools.SliceFor(arr1[:])
	tools.SliceForRange(arr2[0:len(arr2)])

	fmt.Println("------------ 下面展示一维切片 ------------")

	// 1. 基于已有的数组创建切片     arr[0:endIndex] -> arr[:endIndex]    arr[0:len(arr)] -> arr[:]   arr[startIndex:len(arr)] -> arr[startIndex:]
	var slice1 []int = arr1[1:3] 												// 包含左边，不包含右边
	fmt.Println("slice1 = ", slice1)
	fmt.Println("len(slice1) = ", len(slice1)) 									// 2 切片中元素的个数
	fmt.Println("cap(slice1) = ", cap(slice1)) 									// 4 切片容量  容量可以动态变化   通常是元素个数的1.5-2倍
	fmt.Println("&slice1[0] == &arr1[1] ", &slice1[0] == &arr1[1])

	slice1[0] = 200
	fmt.Printf("slice1[0] = %d, arr1[0] = %d \n", slice1[0], arr1[1])

	// 2. 使用make创建切片。底层会自动有一个数组，但是无法直接访问，只能通过切片访问
	var slice2 = make([]string, 2, 4)											// 切片的类型	切片的长度	切片的容量
	slice2[0] = "hello"
	slice2[1] = "world"
	fmt.Println("slice2 = ", slice2)

	// 3. 使用 []数据类型 创建切片
	var slice3 = []string{"你好", "世界"}
	fmt.Println("slice3 = ", slice3)

	// 4. 使用 append() 对原来切片新增内容。创建一个新的切片，先把原来切片的数据拷贝过去，然后再新增
	var slice4 = append(slice3, "hello", "world")
	fmt.Println("slice4 = ", slice4)

	// 5. 使用 copy 对切片进行复制
	var slice5 = make([]string, 10, 20)
	copy(slice5, slice4)
	fmt.Println("slice5 = ", slice5)


	fmt.Println("------------ 下面展示二维数组 ------------")
	var arr5 [2][3]int = [2][3]int{
		{10, 9, 8},
		{7, 6, 5},
	}

	fmt.Println("------------ 下面展示二维切片 ------------")
	var slice6 [][]int = make([][]int, len(arr5))
	for index := range arr5 {												// 多维数组转切片，必须手动转
		slice6[index] = arr5[index][:]
	}
	tools.SliceFor2D(slice6)
	tools.SliceForRange2D(slice6)
}
