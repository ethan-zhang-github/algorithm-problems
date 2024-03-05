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

// CheckInclusion https://leetcode.cn/problems/permutation-in-string/
func CheckInclusion(s1 string, s2 string) bool {
	cache := make(map[byte]int, len(s1))
	for i := 0; i < len(s1); i++ {
		cache[s1[i]]++
	}
	l, r, valid := 0, 0, 0
	win := make(map[byte]int, len(s1))
	for r < len(s2) {
		c := s2[r]
		win[c]++
		r++
		if win[c] == cache[c] {
			valid++
		}
		if r-l > len(s1) {
			if win[s2[l]] == cache[s2[l]] {
				valid--
			}
			win[s2[l]]--
			l++
		}
		if valid == len(cache) {
			return true
		}
	}
	return false
}

// MinWindow https://leetcode.cn/problems/minimum-window-substring/
func MinWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	need := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	ml, mr := 0, 0
	l, r, valid := 0, 0, 0
	win := make(map[byte]int, len(t))
	for r < len(s) {
		win[s[r]]++
		if win[s[r]] == need[s[r]] {
			valid++
		}
		r++
		if valid == len(need) {
			for l < r {
				if n, ok := need[s[l]]; !ok || n < win[s[l]] {
					win[s[l]]--
					l++
				} else {
					break
				}
			}
			if mr == 0 || r-l < mr-ml {
				ml, mr = l, r
			}
		}
	}
	if valid != len(need) {
		return ""
	}
	return s[ml:mr]
}
