package main

import "fmt"

func test1() {
	slice1 := []int{1, 2, 3, 4}

	for _, val := range slice1 {
		if val == 1 {
			val = 5
		}
		// val++
		fmt.Printf("val: %v %p\n", val, &val)
	}

	fmt.Printf("slice1: %v\n", slice1)
}

func test2() {
	slice2 := []int{1, 2, 3, 4}

	for i := range slice2 {
		if slice2[i] == 1 {
			slice2[i] = 5
		}
		// slice2[i]++
	}

	fmt.Printf("slice2: %v\n", slice2)
}

func test3() {
	slice3 := []int{1, 2, 3, 4}

	for i := 0; i < len(slice3); i++ {
		slice3[i]++
	}

	fmt.Printf("slice3: %v\n", slice3)
}

func foo() {
	fmt.Println(*p)
}

var p *int

func main() {
	test1()
	test2()
	// test3()
	// foo()
	// i := 2
	// p = &i
	// foo()
	// fmt.Printf("p: %v\n", p)
}
