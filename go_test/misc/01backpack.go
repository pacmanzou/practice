package main

import (
	"fmt"
)

func main() {
	w := []int{2, 3, 4, 5}
	v := []int{3, 4, 5, 6}
	capacity := 8

	w = append([]int{0}, w...)
	v = append([]int{0}, v...)

	dp := make([]int, 8+1)
	// dp := make([][8 + 1]int, 4+1)

	// for i := 1; i < 4+1; i++ {
	// 	for j := 1; j < capacity+1; j++ {
	// 		if j < w[i] {
	// 			dp[i][j] = dp[i-1][j]
	// 		} else {
	// 			dp[i][j] = max(dp[i-1][j-w[i]]+v[i], dp[i-1][j])
	// 		}
	// 	}
	// }
	// for _, v := range dp {
	// 	fmt.Println(v)
	// }
	// fmt.Println(dp[4][8])
	for i := 1; i < 4+1; i++ {
		for j := capacity; j >= 1; j-- {
			if j >= w[i] {
				dp[j] = max(dp[j-w[i]]+v[i], dp[j])
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(dp[8])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
