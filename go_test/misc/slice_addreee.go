package main

import "fmt"

func main() {
	slice := make([]int, 2, 3)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	fmt.Printf("slice: %v, addr: %p \n", slice, &slice)

	changeSlice(slice)
	fmt.Printf("slice: %v, addr: %p \n", slice, &slice)
	fmt.Println()
	a := []int{1, 2}
	b := a
	a = append(a, 3)
	a[0] = 0
	// b[0] = 0
	fmt.Printf("a: %v addr: %p\n", a, &a)
	fmt.Printf("b: %v addr: %p\n", b, &b)

	// nums := []int{0, 0, 1, 0, 3, 12}
	// // 将切片中的0元素移至末尾,地址发生变化
	// newNums := make([]int, 0, len(nums))
	// temp := 0
	// for _, num := range nums {
	// 	if num == 0 {
	// 		temp++
	// 	} else {
	// 		newNums = append(newNums, num)
	// 	}
	// }
	// // for i := 0; i < len(nums); i++ {
	// // 	if nums[i] == 0 {
	// // 		temp++
	// // 	}else {
	// // 		newNums = append(newNums, nums[i])
	// // 	}
	// // }
	// newNums = append(newNums, make([]int, temp, temp)...)
	// fmt.Printf("nums   : %v %p\n", nums, &nums)
	// fmt.Printf("newNums: %v %p\n", newNums, &newNums)
	// moveZeroes(nums) // 没有拷贝额外的数组
	// fmt.Printf("nums   : %v %p\n", nums, &nums)
}

func changeSlice(s []int) {
	s = append(s, 3)
	s = append(s, 4)
	fmt.Printf("func s: %v, addr: %p \n", s, &s)
	s[1] = 111
	fmt.Printf("func s: %v, addr: %p \n", s, &s)
}

// 在原数组上面操作,不拷贝额外的数组
func moveZeroes(nums []int) {
	// 双指针
	left, right, lenth := 0, 0, len(nums)
	for right < lenth {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
		right++
	}

	// 暴力append
	// i := 0
	// lenth := len(nums)
	// for i < lenth {
	// 	if nums[i] == 0 {
	// 		nums = append(nums[:i], nums[i+1:]...) // 删除等于0的第i项
	// 		nums = append(nums, 0)
	// 		lenth--
	// 	} else {
	// 		i++
	// 	}
	// }
}
