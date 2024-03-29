package array

import (
	"algorithm-problems/array/num_array"
	"sort"
)

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

// CorpFlightBookings https://leetcode.cn/problems/corporate-flight-bookings/
func CorpFlightBookings(bookings [][]int, n int) []int {
	diff := num_array.NewDifference(make([]int, n))
	for _, booking := range bookings {
		diff.Increment(booking[0]-1, booking[1]-1, booking[2])
	}
	return diff.Result()
}

// SearchRange https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func SearchRange(nums []int, target int) []int {
	left := searchLeft(nums, 0, len(nums), target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := searchRight(nums, left-1, len(nums)-1, target)
	return []int{left, right}
}

func searchLeft(nums []int, left int, right int, target int) int {
	if left >= right {
		if left < len(nums) && target == nums[left] {
			return left
		}
		return -1
	}
	mid := left + (right-left)/2
	if target <= nums[mid] {
		return searchLeft(nums, left, mid, target)
	} else {
		return searchLeft(nums, mid+1, right, target)
	}
}

func searchRight(nums []int, left int, right int, target int) int {
	if left >= right {
		if right >= 0 && target == nums[right] {
			return right
		}
		return -1
	}
	mid := right - (right-left)/2
	if target >= nums[mid] {
		return searchRight(nums, mid, right, target)
	} else {
		return searchRight(nums, left, mid-1, target)
	}
}

// Search https://leetcode.cn/problems/binary-search/
func Search(nums []int, target int) int {
	return searchRecursive(nums, target, 0, len(nums)-1)
}

func searchRecursive(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	} else if target > nums[mid] {
		return searchRecursive(nums, target, mid+1, right)
	} else {
		return searchRecursive(nums, target, left, mid-1)
	}
}

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// https://leetcode.cn/problems/3sum/description/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 3, 0, 0)
}

// https://leetcode.cn/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 4, 0, target)
}

func nSum(nums []int, n int, start int, target int) [][]int {
	size := len(nums)
	res := make([][]int, 0)
	if n < 2 || n > size {
		return res
	}
	if n == 2 {
		l, r := start, size-1
		for l < r {
			sum := nums[l] + nums[r]
			lv, rv := nums[l], nums[r]
			if sum > target {
				for l < r && nums[r] == rv {
					r--
				}
			} else if sum < target {
				for l < r && nums[l] == lv {
					l++
				}
			} else {
				res = append(res, []int{lv, rv})
				for l < r && nums[r] == rv {
					r--
				}
				for l < r && nums[l] == lv {
					l++
				}
			}
		}
		return res
	}
	for i := start; i < size; i++ {
		sub := nSum(nums, n-1, i+1, target-nums[i])
		for _, v := range sub {
			res = append(res, append(v, nums[i]))
		}
		for i < size-1 && nums[i+1] == nums[i] {
			i++
		}
	}
	return res
}
