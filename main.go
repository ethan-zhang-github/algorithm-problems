package main

import (
	List "algorithm-problems/linked_list"
	"fmt"
)

func main() {
	list := List.Cons(1, 4, 3, 2, 5, 2)

	res := List.Partition(list, 3)

	fmt.Println(res)
}
