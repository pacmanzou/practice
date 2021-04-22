package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 7, 2, 9, 1, 77}
	fmt.Println(a)
	fmt.Println()
	// 优化取中位数
	quickSort(a, 0, len(a)-1)

	// 单向
	// quickSort1(a, 0, len(a)-1)

	// 严魏敏 填坑
	// quickSort2(a, 0, len(a)-1)

	// 堆
	// heapSort(a, 0, len(a))

	// 冒泡 标记
	// bubbleSort(a)

	// 插入
	// insertSort(a)

	fmt.Println()
	fmt.Println(a)
}

// 优化取中位数
// [5 4 7 2 9 1 77]

// [2 4 1 5 9 7 77]
// [1 2 4 5 9 7 77]
// [1 2 4 5 7 9 77]
func doPivot(s []int, left, right int) int {
	mid := int(uint(left+right) >> 1)
	pivot := methodThtee(s, mid, left, right)
	for {
		for ; left < right && s[right] > pivot; right-- {
		}
		for ; left < right && s[left] < pivot; left++ {
		}
		if left >= right {
			break
		}
		s[left], s[right] = s[right], s[left]
	}
	fmt.Println(s)
	return left
}

func methodThtee(s []int, mid, left, right int) int {
	if s[left] > s[mid] {
		s[left], s[mid] = s[mid], s[left]
	}
	if s[left] > s[right] {
		s[left], s[right] = s[right], s[left]
	}
	if s[mid] > s[right] {
		s[mid], s[right] = s[right], s[mid]
	}
	// s[left]<=s[mid]<=s[right]
	return s[mid]
}

func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}
	pivot := doPivot(s, left, right)
	quickSort(s, left, pivot-1)
	quickSort(s, pivot+1, right)
}

// 单向
// [5 4 7 2 9 1 77]

// [1 4 2 5 9 7 77]
// [1 4 2 5 9 7 77]
// [1 2 4 5 9 7 77]
// [1 2 4 5 7 9 77]
func doPivot1(s []int, left, right int) int {
	pivot := s[left]
	mark := left
	cur := left + 1
	for cur < right+1 {
		if s[cur] < pivot {
			mark++
			s[cur], s[mark] = s[mark], s[cur]
		}
		cur++
	}
	s[left], s[mark] = s[mark], s[left]
	fmt.Println(s)
	return mark
}

func quickSort1(s []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := doPivot2(s, left, right)
	quickSort1(s, left, pivotIndex-1)
	quickSort1(s, pivotIndex+1, right)
}

//严魏敏版 填坑
// 5, 4, 7, 2, 9, 1, 77

// 1, 4, 7, 2, 9, 5, 77
// 1, 4, 5, 2, 9, 7, 77
// 1, 4, 5, 2, 9, 7, 77
// 1, 4, 2, 5, 9, 7, 77

func doPivot2(s []int, left, right int) int {
	pivot := s[left]
	for {
		for ; left < right && s[right] > pivot; right-- {
		}
		s[left] = s[right]
		for ; left < right && s[left] < pivot; left++ {
		}
		s[right] = s[left]
		if left >= right {
			break
		}
	}
	s[left] = pivot
	fmt.Println(s)
	return left
}

func quickSort2(s []int, left, right int) {
	if left >= right {
		return
	}
	pivot := doPivot(s, left, right)
	quickSort2(s, left, pivot-1)
	quickSort2(s, pivot+1, right)
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
