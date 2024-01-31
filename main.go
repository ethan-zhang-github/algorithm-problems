package main

import (
	Array "algorithm-problems/array"
	"fmt"
)

func main() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(Array.RemoveDuplicates(a))
	fmt.Println(a)
}
