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

func TestCheckInclusion(t *testing.T) {
	fmt.Println(CheckInclusion("abcdxabcde", "abcdeabcdx"))
}

func TestMinWindow(t *testing.T) {
	fmt.Println(MinWindow("ab", "a"))
}
