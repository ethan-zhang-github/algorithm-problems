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
