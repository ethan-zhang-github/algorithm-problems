package main

import (
	List "algorithm-problems/linked_list"
	"fmt"
)

func main() {
	list := List.Cons(1, 4, 3, 2, 5, 2, 7)

	fmt.Println(List.ReverseBetween(list, 3, 5))

	list2 := List.Cons(1, 4, 3, 2, 5, 2, 7)
	fmt.Println(List.ReverseN(list2, 3))
}
