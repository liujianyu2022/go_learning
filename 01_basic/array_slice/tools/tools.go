package tools

import "fmt"

// 注意：长度属于数组类型的一部分。比如 [3]int   [5]int 就是不同的类型
// 下面只能处理 [5]int 类型的数组
func ForRange1(arr [5]int) {
	for index, value := range arr {
		fmt.Printf("第%d个的结果为：%d \n", index, value)
	}
}

func For1(arr [5]int){
	for i := 0; i < len(arr); i++ {
		fmt.Printf("第%d个的结果为：%d \n", i, arr[i])
	} 
}


func ForRange2(arr [3]string){
	for index, value := range arr {
		fmt.Printf("第%d个的结果为：%s \n", index, value)
	}
}

func For2(arr [3]string){
	for i := 0; i < len(arr); i++ {
		fmt.Printf("第%d个的结果为：%s \n", i, arr[i])
	}
}

// 由于 [5]int, [3]int 是两个不同的数据类型，遍历函数会很多
// 因此最佳实践是转为切片
func SliceForRange(slice []int){
	for index, value := range slice {
		fmt.Printf("slice[%d] = %d \t", index, value)
	}
	fmt.Println()
}
func SliceFor(slice []int){
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] = %d \t", i, slice[i])
	}
	fmt.Println()
}


// 遍历多维数组
func SliceFor2D(arr [][]int){
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++{
			fmt.Printf("arr[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}
}
func SliceForRange2D(arr [][]int){
	for i, row := range arr {
		for j, value := range row {
			fmt.Printf("arr[%d][%d] = %d\n", i, j, value)
		}
	}
}