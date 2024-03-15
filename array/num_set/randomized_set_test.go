package num_set

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	randomizedSet := Constructor()
	randomizedSet.Insert(0)
	randomizedSet.Insert(1)
	randomizedSet.Remove(0)
	randomizedSet.Insert(2)
	randomizedSet.Remove(1)
	fmt.Println(randomizedSet.GetRandom())
}
