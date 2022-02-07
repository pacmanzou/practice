package main

import "fmt"

func main() {
	ch := make(chan struct{})
	go func(a, b int) {
		fmt.Println(a + b)
		<-ch
	}(3, 6)
	ch <- struct{}{}
	fmt.Println("hello world")
}
