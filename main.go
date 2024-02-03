package main

import (
	Array "algorithm-problems/array"
	"fmt"
)

func main() {
	a := []int{2, 2, 3}
	fmt.Println(Array.RemoveElement(a, 2))
	fmt.Println(a)
}
