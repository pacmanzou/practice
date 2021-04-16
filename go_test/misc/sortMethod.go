package main

// 冒泡
func bubbleSort(nums []int) {
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
