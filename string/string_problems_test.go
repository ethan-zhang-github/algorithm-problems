package string

import (
	"fmt"
	"testing"
)

func TestReverseMessage(t *testing.T) {
	s := "the sky is blue"
	res := ReverseMessage(s)
	fmt.Println(res)
}
