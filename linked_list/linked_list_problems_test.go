package linked_list

import (
	"fmt"
	"testing"
)

func TestIsPalindromeRecursive(t *testing.T) {
	l := Cons(1, 2, 2, 1)
	fmt.Println(isPalindromeRecursive(l))
}
