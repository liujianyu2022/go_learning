package main

import "fmt"

// 在 runtime/string.go 中，字符串的定义如以下：

// type stringStruct struct {		// 在 go 中，字符串的本质是一个结构体
// 	str unsafe.Pointer				// str 指针指向底层的 Byte数组
// 	len int							// len 返回的是字节数组的长度，不是字符个数
// }

// 在 runtime/slice.go 中，切片的定义如下：

// type slice struct {
// 	array unsafe.Pointer			// 切片的本质是对数组的引用
// 	len   int						// 切片所引用的byte数组的长度
// 	cap   int						// 数组底层byte数组的长度
// }

func main() {
	str := "你好hello"

	fmt.Println("the length of byte array is ", len(str)) // 11 = 2 * 3 + 5 * 1

	// 下面这种遍历方式是错误的
	// 在底层的byte数组中，一个汉字是3个字节表示的
	// len(str) 返回的是底层byte数组的长度，而不是字符个数。因此会导致乱码
	// Go 的 string 是 UTF-8 编码的字节序列, 非 ASCII 字符（如中文）在 UTF-8 中可能占用 1~4 个字节，这是一种变长的编码方式
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	fmt.Println("first byte is ", str[0]) // 直接使用索引 str[0]，获取的是第一个字节（UTF-8 编码下的首字节），而非完整字符，出现乱码

	runes := []rune(str) // 将字符串转换为 []rune（Unicode 码点切片），识别完整的 Unicode 字符，然后按索引访问
	fmt.Println("first char is ", string(runes[0]))

	// 使用 range 遍历的时候，会对字符进行解码，循环时自动按 Unicode 字符迭代（而非字节）
	for _, char := range str {
		fmt.Printf("%c\n", char)
	}

	fmt.Println("-------------------------------------")

	var arr [4]string = [4]string{"你", "好", "hello", "world"}
	slice1 := arr[1:3]
	fmt.Println("length of slice1 = ", len(slice1))   			// 2 当前元素个数
	fmt.Println("capacity of slice1 = ", cap(slice1)) 			// 3 容量会从startIndex开始，一直到这个数组末尾。

	slice2 := []string{"你", "好", "hello"}             		// 字面量创建的切片
	fmt.Println("length of slice2 = ", len(slice2))   		    // 3 当前元素个数
	fmt.Println("capacity of slice2 = ", cap(slice2)) 			// 3 字面量创建的切片，初始容量等于长度（除非显式指定容量）。

	slice2 = append(slice2, "world")
	fmt.Println("length of slice2 = ", len(slice2))   			// 4 当前元素个数
	fmt.Println("capacity of slice2 = ", cap(slice2)) 			// 6 切片扩容了。当旧容量 < 1024 时，容量通常按2倍增长；当超过时，按照25%增长

	slice3 := make([]string, 3, 5)
	fmt.Println("length of slice3 = ", len(slice3))
	fmt.Println("capacity of slice3 = ", cap(slice3))

	fmt.Println("--------------------------------------")

	bytes := []byte(str) // 字符串 → []byte（字节切片）
	runes = []rune(str)  // 字符串 → []rune（Unicode 码点切片）

	fmt.Println("bytes = ", bytes) // [228 189 160 229 165 189 104 101 108 108 111]
	fmt.Println("runes = ", runes) // [20320 22909 104 101 108 108 111]
}
