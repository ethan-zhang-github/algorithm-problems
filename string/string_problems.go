package string

import (
	"unicode"
)

// ReverseMessage https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/
func ReverseMessage(s string) string {
	i := len(s) - 1
	for i >= 0 && unicode.IsSpace(rune(s[i])) {
		i--
	}
	j := i
	res := ""
	for i >= 0 {
		for i >= 0 && !unicode.IsSpace(rune(s[i])) {
			i--
		}
		res += s[i+1:j+1] + " "
		for i >= 0 && unicode.IsSpace(rune(s[i])) {
			i--
		}
		j = i
	}
	return res[0 : len(res)-1]
}

// LengthOfLongestSubstring https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	maxLen, l, r := 0, 0, 0
	win := make(map[rune]int)
	for l < len(s) && r < len(s) {
		c := rune(s[r])
		if i, e := win[c]; e && i >= l {
			l = i + 1
		}
		win[c] = r
		r++
		if r-l > maxLen {
			maxLen = r - l
		}
	}
	return maxLen
}
