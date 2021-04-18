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
	return s[i] > s[j]
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
	nums := []int{3, 2, 1, 5, 6, 4, 8}
	n := &IntHeap{}
	for _, num := range nums {
		*n = append(*n, num)
	}
	heap.Init(n)

	// 改变了第i个元素时，用fix修复堆，恢复堆结构
	// (*n)[1] = 7
	// heap.Fix(n, 1)

	fmt.Println(n)
	length := len(*n)
	for i := 0; i < length; i++ {
		fmt.Println(heap.Pop(n))
		// 和heap.Pop(n)结果不一样
		// fmt.Println(n.Pop())
	}
}
