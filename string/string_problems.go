package string

import "unicode"

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
