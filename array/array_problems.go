package array

import "algorithm-problems/array/num_array"

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

// ReverseString https://leetcode.cn/problems/reverse-string/
func ReverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// LongestPalindrome https://leetcode.cn/problems/longest-palindromic-substring/
func LongestPalindrome(s string) string {
	maxL, maxR, maxLen := 0, 1, 1
	for i := 0; i < len(s); i++ {
		l, r := palindrome(s, i, i)
		if r-l > maxLen {
			maxL, maxR, maxLen = l, r, r-l
		}
		l, r = palindrome(s, i, i+1)
		if r-l > maxLen {
			maxL, maxR, maxLen = l, r, r-l
		}
	}
	return s[maxL:maxR]
}

func palindrome(s string, l int, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r
}

// CarPooling https://leetcode.cn/problems/car-pooling/
func CarPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	diff := num_array.NewDifference(nums)
	for _, trip := range trips {
		diff.Increment(trip[1], trip[2]-1, trip[0])
	}
	res := diff.Result()
	for _, val := range res {
		if val > capacity {
			return false
		}
	}
	return true
}
