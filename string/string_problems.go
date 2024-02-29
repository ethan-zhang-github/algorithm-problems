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

// FindAnagrams https://leetcode.cn/problems/find-all-anagrams-in-a-string/
func FindAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	pMap := make(map[rune]int, len(p))
	for _, c := range p {
		pMap[c]++
	}
	sMap := make(map[rune]int, len(p))
	l, r := 0, len(p)-1
	for i := l; i <= r; i++ {
		sMap[rune(s[i])]++
	}
	res := make([]int, 0)
	for r < len(s) {
		if equals(pMap, sMap) {
			res = append(res, l)
		}
		if r == len(s)-1 {
			break
		}
		sMap[rune(s[l])]--
		l++
		r++
		sMap[rune(s[r])]++
	}
	return res
}

func equals(map1 map[rune]int, map2 map[rune]int) bool {
	for k1, v1 := range map1 {
		v2, ok := map2[k1]
		if v1 > 0 && (!ok || v1 != v2) {
			return false
		}
	}
	return true
}
