package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 7, 2, 9, 1, 77, 2, 5, 23, 23, 23, 199, 45, 23, 5, 999, 44, 22, 1, 3, 1, 25, 88}
	// a := []int{1, 2, 3, 4, 5}
	// a := []int{5, 4, 3, 2, 1}
	fmt.Printf("%v\n\n", a)

	// 冒泡
	// bubbleSort(a)

	// 插入
	// insertSort(a)

	// 希尔
	// shellSort(a)

	// 选择
	// selectSort(a)

	// 快速
	// quickSort(a, 0, len(a)-1)

	// 归并
	// mergeSort(a, 0, len(a))
	// fmt.Println(a)

	// 堆
	// heapSort(a, 0, len(a))
}

// 冒泡
func bubbleSort(nums []int) {
	// 记录遍历是否交换了元素位置
	isChange := false
	// countF := 0
	// countS := 0
	for i := 0; i < len(nums)-1; i++ {
		// countF++
		// fmt.Printf("first: %v  i = %v\n", nums, i)
		for j := 0; j < len(nums)-1-i; j++ {
			// countS++
			// fmt.Printf("second: %v  j = %v\n", nums, j)
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
	// fmt.Printf("\n%v\n", nums)
	// fmt.Printf("\ncountF = %v, countS = %v", countF, countS)
}

// 插入
func insertSort(nums []int) {
	// countF, countS := 0, 0
	for i := 1; i < len(nums); i++ {
		// countF++
		// fmt.Printf("first: %v  i = %v\n", nums, i)
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			// countS++
			// fmt.Printf("second: %v  j = %v\n", nums, j)
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	// fmt.Printf("\n%v\n", nums)
	// fmt.Printf("\ncountF = %v, countS = %v", countF, countS)
}

// 希尔
func shellSort(nums []int) {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			if nums[i] < nums[i-gap] {
				nums[i], nums[i-gap] = nums[i-gap], nums[i]
			}
		}
	}
	insertSort(nums)
	fmt.Printf("nums: %v\n", nums)
}

// 选择
func selectSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] <= nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	fmt.Println(nums)
}

// 快排
func doPivot(s []int, left, right int) int {
	mid := int(uint(left+right) >> 1)
	methodThree(s, left, mid, right)
	s[left], s[mid] = s[mid], s[left]

	pivot := left
	for {
		for ; left < right && s[right] >= s[pivot]; right-- {
		}
		for ; left < right && s[left] <= s[pivot]; left++ {
		}
		if left >= right {
			break
		}
		s[left], s[right] = s[right], s[left]
	}
	s[pivot], s[left] = s[left], s[pivot]
	fmt.Println(s)
	return left
}

func methodThree(s []int, left, mid, right int) {
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
}

func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}
	pivot := doPivot(s, left, right)
	quickSort(s, left, pivot-1)
	quickSort(s, pivot+1, right)
}

// 归并
func merge(nums []int, left, mid, right int) {
	leftSize := mid - left
	rightSize := right - mid
	newSize := leftSize + rightSize
	res := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := nums[left+l]
		rValue := nums[mid+r]

		if lValue < rValue {
			res = append(res, lValue)
			l++
		} else {
			res = append(res, rValue)
			r++
		}
	}
	res = append(res, nums[left+l:mid]...)
	res = append(res, nums[mid+r:right]...)
	for i := 0; i < newSize; i++ {
		nums[left+i] = res[i]
	}
	// return
}

func mergeSort(nums []int, left, right int) {
	if right-left <= 1 {
		return
	}
	mid := int(uint(left+right) >> 1)
	mergeSort(nums, left, mid)
	mergeSort(nums, mid, right)
	merge(nums, left, mid, right)
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
