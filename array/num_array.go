package array

type NumArray struct {
	SumArray []int
}

func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums))
	sumArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sumArray[i] = sumArray[i-1] + nums[i]
	}
	return NumArray{SumArray: sumArray}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.SumArray[right]
	}
	return this.SumArray[right] - this.SumArray[left-1]
}
