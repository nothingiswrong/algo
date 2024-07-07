package main

import (
	"fmt"
)

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}


func main() {
	// var x rune = 'A'
	// var y rune = 'B'
	var s1 string = "ABCD"
	var s2 string = "ABCD"
	fmt.Println(s1 < s2)
}
