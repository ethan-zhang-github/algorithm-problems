package num_array

import (
	"fmt"
	"testing"
)

func TestAdvantageCount(t *testing.T) {
	nums1 := []int{12, 24, 8, 32}
	nums2 := []int{13, 25, 32, 11}
	res := AdvantageCount(nums1, nums2)
	fmt.Println(res)
}
