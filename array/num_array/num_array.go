package num_array

// NumArray https://leetcode.cn/problems/range-sum-query-immutable/description/
type NumArray struct {
	sumArray []int
}

func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums))
	sumArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sumArray[i] = sumArray[i-1] + nums[i]
	}
	return NumArray{sumArray: sumArray}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sumArray[right]
	}
	return this.sumArray[right] - this.sumArray[left-1]
}
