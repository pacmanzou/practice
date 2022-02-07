package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 7, 2, 9, 1, 77, 2, 5, 23, 23, 23, 199, 45, 23, 5, 999, 44}
	fmt.Println(a)
	fmt.Println()

	// quickSort(a, 0, len(a)-1)

	// 优化取中位数
	// quickSortMedian(a, 0, len(a)-1)

	// 堆
	// heapSort(a, 0, len(a))

	// 冒泡 标记
	bubbleSort(a)

	// 插入
	// insertSort(a)
}

func doPivot(s []int, left, right int) int {
	pivot := s[left]
	i, j := left, right
	for {
		for ; i < j && s[j] >= pivot; j-- {
		}
		for ; i < j && s[i] <= pivot; i++ {
		}
		if i >= j {
			break
		}
		s[i], s[j] = s[j], s[i]
	}
	s[left], s[i] = s[i], s[left]
	fmt.Println(s)
	return i
}

func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}
	pivot := doPivot(s, left, right)
	quickSort(s, left, pivot-1)
	quickSort(s, pivot+1, right)
}

// 优化取中位数
func doPivotMedian(s []int, left, right int) int {
	pivot := methodThree(s, left, right)
	i, j := left, right
	for {
		for ; i < j && s[j] >= pivot; j-- {
		}
		for ; i < j && s[i] <= pivot; i++ {
		}
		if i >= j {
			break
		}
		s[i], s[j] = s[j], s[i]
	}
	s[left], s[i] = s[i], s[left]
	fmt.Println(s)
	return i
}

func methodThree(s []int, left, right int) int {
	mid := int(uint(left+right) >> 1)

	// s[left]<=s[mid]<=s[right]
	if s[left] > s[mid] {
		s[left], s[mid] = s[mid], s[left]
	}
	if s[left] > s[right] {
		s[left], s[right] = s[right], s[left]
	}
	if s[mid] > s[right] {
		s[mid], s[right] = s[right], s[mid]
	}

	// let's move the middle number to the left
	s[left], s[mid] = s[mid], s[left]

	return s[left]
}

func quickSortMedian(s []int, left, right int) {
	if left >= right {
		return
	}
	pivot := doPivotMedian(s, left, right)
	quickSortMedian(s, left, pivot-1)
	quickSortMedian(s, pivot+1, right)
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
		fmt.Println(data)
	}
}

// 冒泡
func bubbleSort(nums []int) {
	// 标记出最后一次交换元素的位置
	flag := true
	for i := 0; i < len(nums)-1 && flag; i++ {
		flag = false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		fmt.Println(nums)
	}
}

// 插入
func insertSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
		fmt.Println(nums)
	}
}
