package array

// TwoSum https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
func TwoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}

// RemoveDuplicates https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}

// RemoveElement https://leetcode.cn/problems/remove-element/
func RemoveElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}

// MoveZeroes https://leetcode.cn/problems/move-zeroes/
func MoveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
