package main

import "fmt"

func testPoint() {
	s := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for i, v := range s {
		m[i] = &v
	}
	printMapPoint(m)
}

func test() {
	s := []int{0, 1, 2, 3, 4}
	m := make(map[int]int)

	for i := 0; i < len(s); i++ {
		m[i] = s[i]
	}
	printMap(m)
}

func printMapPoint(m map[int]*int) {
	for k, v := range m {
		fmt.Printf("map[%v]=%v\n", k, *v)
	}
}

func printMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("map[%v]=%v\n", k, v)
	}
}

func main() {
	fmt.Println("test:")
	test()
	fmt.Println()
	fmt.Println("testPoint:")
	testPoint()
}
