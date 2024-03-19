package tree

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	tree := Constructor([]int{3, 9, 20, NULL, NULL, 15, 7})
	fmt.Println(tree)
}
