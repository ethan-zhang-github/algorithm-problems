package array

import (
	"fmt"
	"testing"
	"unicode"
)

func TestCarPooling(t *testing.T) {
	trips := [][]int{{2, 1, 5}, {3, 5, 7}}
	res := CarPooling(trips, 3)
	fmt.Println(res)
}

func TestCorpFlightBookings(t *testing.T) {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	res := CorpFlightBookings(bookings, 5)
	fmt.Println(res)
}

func TestReverseWords(t *testing.T) {
	s := " a"

	for _, c := range s {
		fmt.Println(unicode.IsSpace(c))
	}
}

func TestSearchRange(t *testing.T) {
	//fmt.Println(SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(SearchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(SearchRange([]int{2, 2}, 3))
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
