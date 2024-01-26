package main

import (
	List "algorithm-problems/linked_list"
	"fmt"
)

func main() {
	list2 := List.Cons(1, 4, 3, 2, 5, 2, 7)
	fmt.Println(List.ReverseKGroup(list2, 10))
}
