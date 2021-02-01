package main

import "fmt"

type Slice struct {
	data []int
}

func (s Slice) addValue(x int) {
	s.data = append(s.data, x)
}
func (s *Slice) addPointer(x int) {
	s.data = append(s.data, x)
}

func main() {
	// 因为slice是引用类型，所以做了如下测试
	// method
	fmt.Println("method")
	a := Slice{}
	a.addValue(2)
	a.addValue(3)
	fmt.Println(a)
	a.addPointer(4)
	a.addPointer(5)
	fmt.Println(a)
	fmt.Println()

	// func
	fmt.Println("func")
	b := []int{}
	addValue(b, 2, 3)
	fmt.Println(b)
	addPointer(&b, 4, 5)
	fmt.Println(b)
	fmt.Println()

	// 当函数不改变切片长度时，也就是不改变其地址，此时参数不必传入指针就能改变其含有的值
	// 函数内有append，如果在初始化容量以内操作时，可以不用传Pointer,反之一定要传入
	fmt.Println("func")
	c := []int{1, 2, 3, 4}
	changeV(c)
	fmt.Println(c)
	changeP(&c)
	fmt.Println(c)
}

func changeV(x []int) {
	x[1] = 0
}

func changeP(x *[]int) {
	(*x)[2] = 0
}

func addValue(x []int, v ...int) {
	x = append(x, v...)
}

func addPointer(x *[]int, v ...int) {
	*x = append(*x, v...)
}
