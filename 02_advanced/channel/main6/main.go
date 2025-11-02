package main

import (
	"fmt"
	"strings"
)

// 管道阶段函数类型
type Stage func(<-chan string) <-chan string

func main() {
	// 创建管道阶段
	upperStage := createStage(toUpperCase)
	suffixStage := createStage(addSuffix)

	// 构建管道：输入 → 转大写 → 加后缀 → 输出
	input := generateInput("hello", "world")
	output := suffixStage(upperStage(input))

	// 收集结果
	fmt.Println("============== 处理结果： ==============")
	for result := range output {
		fmt.Println("最终结果:", result)
	}
	fmt.Println("所有数据处理完成")
}

// 创建管道阶段
func createStage(processor func(string) string) Stage {
	return func(input <-chan string) <-chan string {
		output := make(chan string)
		go func() {
			defer close(output)
			for data := range input {
				result := processor(data)
				output <- result
			}
		}()
		return output
	}
}

// 生成输入数据
func generateInput(data ...string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, d := range data {
			ch <- d
		}
	}()
	return ch
}

// 处理函数
func toUpperCase(data string) string {
	result := strings.ToUpper(data)
	fmt.Printf("toUpperCase: %s -> %s\n", data, result)
	return result
}

func addSuffix(data string) string {
	result := data + " !"
	fmt.Printf("addSuffix: %s -> %s\n", data, result)
	return result
}