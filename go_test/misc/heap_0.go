package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (s IntHeap) Len() int {
	return len(s)
}
func (s IntHeap) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s IntHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s *IntHeap) Pop() interface{} {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *IntHeap) Push(x interface{}) {
	*s = append(*s, x.(int))
}
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	n := IntHeap(nums)
	// 当实现接口的新类型里面有指针方法时(比如上面的Pop和Push)，不能把值对象赋值给该接口类型的变量
	// cannot use n (variable of type TopHeap) as heap.Interface value in argument to heap.Init: missing method Pop (Pop has pointer receiver)
	// heap.Init(n)
	fmt.Println(n)
	heap.Init(&n)
	fmt.Println(n)

	// a := heap.Pop(&n) // Pop出去的是interface{}空接口
	length := len(n)
	for i := 0; i < length; i++ {
		fmt.Println(heap.Pop(&n))
	}
	// m := IntHeap{3, 2, 1, 5, 6, 4}
	// fmt.Println(m)
	// heap.Init(&m)
	// fmt.Println(m)

	// // // push生成的heap不需要init
	// s := IntHeap{}
	// heap.Push(&s, 3)
	// heap.Push(&s, 2)
	// heap.Push(&s, 1)
	// heap.Push(&s, 4)
	// fmt.Println(heap.Pop(&s))
	// fmt.Println(heap.Pop(&s))
	// fmt.Println(heap.Pop(&s))
	// fmt.Println(heap.Pop(&s))
}
