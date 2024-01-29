package main

import (
	List "algorithm-problems/linked_list"
	"fmt"
)

func main() {
	list2 := List.Cons(1, 2)
	fmt.Println(List.IsPalindrome(list2))
}
