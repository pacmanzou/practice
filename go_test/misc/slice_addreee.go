package main

import "fmt"

func main() {

	nums := []int{0, 0, 1, 3, 12}
	a := nums[1:]
	b := nums[:2]
	c := nums[1:3]
	fmt.Printf("原nums切片:\t%d\t地址:\t%p\t容量:\t%d\n", nums, nums, cap(nums))
	fmt.Printf("a 截取后面:\t%d\t地址:\t%p\t容量:\t%d\n", a, a, cap(a))
	fmt.Printf("b 截取前面:\t%d\t\t地址:\t%p\t容量:\t%d\n", b, b, cap(b))
	// a[s:e]和a[s:]的cap和地址均相等
	fmt.Printf("c 截取中间:\t%d\t\t地址:\t%p\t容量:\t%d\n", c, c, cap(c))
	a = append(a, 9)
	fmt.Println("\nappend a 9")
	fmt.Printf("原nums切片:\t%d\t地址:\t%p\t容量:\t%d\n", nums, nums, cap(nums))
	fmt.Printf("a 截取后面:\t%d\t地址:\t%p\t容量:\t%d\n", a, a, cap(a))
	fmt.Printf("b 截取前面:\t%d\t\t地址:\t%p\t容量:\t%d\n", b, b, cap(b))
	fmt.Printf("c 截取中间:\t%d\t\t地址:\t%p\t容量:\t%d\n", c, c, cap(c))
	b = append(b, 8)
	fmt.Println("\nappend b 8")
	fmt.Printf("原nums切片:\t%d\t地址:\t%p\t容量:\t%d\n", nums, nums, cap(nums))
	fmt.Printf("a 截取后面:\t%d\t地址:\t%p\t容量:\t%d\n", a, a, cap(a))
	fmt.Printf("b 截取前面:\t%d\t\t地址:\t%p\t容量:\t%d\n", b, b, cap(b))
	fmt.Printf("c 截取中间:\t%d\t\t地址:\t%p\t容量:\t%d\n", c, c, cap(c))
	fmt.Println()

	// 将切片中的0元素移至末尾,地址发生变化
	newNums := make([]int, 0, len(nums))
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			temp++
		} else {
			newNums = append(newNums, nums[i])
		}
	}
	for i := 0; i < temp; i++ {
		newNums = append(newNums, 0)
	}
	fmt.Printf("原切片:\t%d %p %d\n", nums, nums, cap(nums))
	fmt.Printf("新切片1:\t%d %p %d\n", newNums, newNums, cap(newNums))
	moveZeroes(nums) // 没有拷贝额外的数组
	fmt.Printf("新切片2:\t%d %p %d\n", nums, nums, cap(nums))
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
