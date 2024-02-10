package array

import (
	"fmt"
	"testing"
)

func TestCarPooling(t *testing.T) {
	trips := [][]int{{2, 1, 5}, {3, 5, 7}}
	res := CarPooling(trips, 3)
	fmt.Println(res)
}
