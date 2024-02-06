package main

import (
	Array "algorithm-problems/array"
	"algorithm-problems/linked_list"
	"fmt"
)

func main() {
	res := Array.LongestPalindrome("cbbd")
	fmt.Println(res)

	list := linked_list.Cons(1, 1, 2, 3, 3)
	linked_list.DeleteDuplicates(list)
	fmt.Println(list)
}
