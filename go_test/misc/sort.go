package main

import (
	"fmt"
	"sort"
)

type ByteSlice []byte

func (s ByteSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s ByteSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByteSlice) Len() int {
	return len(s)
}

type RuneSlice []rune

func (s RuneSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s RuneSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s RuneSlice) Len() int {
	return len(s)
}

func main() {
	// []int
	intSlice := []int{3, 4, 2, 1}
	sort.Ints(intSlice)
	// sort.Sort(sort.IntSlice(intSlice))
	fmt.Println(intSlice)
	// reverse
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice)

	// sort.Search
	// x := 3
	// pos := sort.Search(len(intSlice), func(i int) bool {
	// 	return intSlice[i] >= x
	// })
	// if pos < len(intSlice) && intSlice[pos] == x {
	// 	fmt.Println(x, "在i中的位置是", pos)
	// } else {
	// 	fmt.Println("intSlice不包含元素", x)
	// }

	// []float64
	// f := []float64{2.3, 4.0, 1.23, 5}
	// sort.Float64s(f)
	// // sort.Sort(sort.Float64Slice(f))
	// fmt.Println(f) // 1.23, 2.3, 4, 5
	// fmt.Println(sort.SearchFloat64s(f, 2.3))

	// []string
	// s := []string{"bf", "a", "bc", "s"}
	// sort.Strings(s)
	// // sort.Sort(sort.StringSlice(s))
	// fmt.Println(s)

	// []byte
	// b := []byte{'s', 'b', 'a', 'c'}
	// sort.Sort(ByteSlice(b))
	// fmt.Printf("%c\n", b)

	// []rune
	// r := []rune{'b', 'a', 'c', 'n'}
	// sort.Sort(RuneSlice(r))
	// fmt.Printf("%c\n", r)
}

// sort.Search
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
