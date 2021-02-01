package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	key   string
	value int
}

type maxHeap []node

func (s maxHeap) Len() int {
	return len(s)
}

func (s maxHeap) Less(i, j int) bool {
	if s[i].value == s[j].value {
		return s[i].key < s[j].key
	}
	return s[i].value > s[j].value
}

func (s maxHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *maxHeap) Pop() interface{} {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *maxHeap) Push(x interface{}) {
	*s = append(*s, x.(node))
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	a := make(map[string]int, len(words))
	var nodeMaxHeap maxHeap
	for i := 0; i < len(words); i++ {
		a[words[i]]++
	}
	for k, v := range a {
		heap.Push(&nodeMaxHeap, node{
			key:   k,
			value: v,
		})
	}
	k := 3
	res := make([]string, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&nodeMaxHeap).(node).key)
	}
	fmt.Println(res)
}
