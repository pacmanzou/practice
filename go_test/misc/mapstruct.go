package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]*Math{"foo": &Math{1, 2}}

var n = map[string]Math{"foo": Math{1, 2}}

func main() {
	m["foo"].x = 3
	fmt.Printf("m[\"foo\"]: %v\n", m["foo"])

	tmp := n["foo"]
	tmp.x = 3
	n["foo"] = tmp
	fmt.Printf("n[\"foo\"]: %v\n", n["foo"])
}
