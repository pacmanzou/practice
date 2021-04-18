package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 1, 5, 6, 4}
	heapSort(a, 0, len(a))
	// bubbleSort(a)
	// insertSort(a)
	fmt.Println(a)
}

// 堆
func siftDown(data []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		// 大顶堆 升序
		if child+1 < hi && data[child] < data[child+1] {
			child++
		}
		if data[root] > data[child] {
			return
		}
		// 小顶堆 降序
		// if child+1 < hi && data[child] > data[child+1] {
		// 	child++
		// }
		// if data[root] < data[child] {
		// 	return
		// }
		data[root], data[child] = data[child], data[root]
		root = child
	}
}

func heapSort(data []int, a, b int) {
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi)
	}

	// Pop elements, largest , into end of data.
	for i := hi - 1; i >= 0; i-- {
		data[lo], data[i] = data[i], data[lo]
		siftDown(data, lo, i)
	}
}

// 冒泡
func bubbleSort(nums []int) {
	// flag标记优化
	flag := true
	for i := 0; i < len(nums)-1 && flag; i++ {
		flag = false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
	}
}

// 插入
func insertSort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}
