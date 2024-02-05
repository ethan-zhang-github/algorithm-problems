package main

import (
	Array "algorithm-problems/array"
	"fmt"
)

func main() {
	a := []int{0, 1, 0, 3, 12}
	Array.MoveZeroes(a)
	fmt.Println(a)
}
