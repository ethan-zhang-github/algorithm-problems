package string

import (
	"fmt"
	"testing"
)

func TestReverseMessage(t *testing.T) {
	s := "the sky is blue"
	res := ReverseMessage(s)
	fmt.Println(res)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(LengthOfLongestSubstring("abcabcbb"))
}

func TestFindAnagrams(t *testing.T) {
	fmt.Println(FindAnagrams("cbaebabacd", "abc"))
	fmt.Println(FindAnagrams("abab", "ab"))
}
