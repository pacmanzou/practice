package main

import "fmt"

func main() {
	// var add func(a, b int)
	add := func(a, b int) int {
		// 内置函数如果里面调用了自己，那么就要var声明
		// add(3, 5)
		return a + b
	}
	fmt.Println(add(3, 6))
}
