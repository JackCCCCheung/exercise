package top75

import (
	"fmt"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	println("---------" + GcdOfStrings("abcabc", "abc"))
}

func TestKidsWithCandies(t *testing.T) {
	fmt.Printf("---------%v", kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func TestCanPlaceFlowers(t *testing.T) {
	fmt.Printf("-------%v\n", canPlaceFlowers([]int{0, 1, 0}, 1))
}

func TestReverseVowels(t *testing.T) {
	fmt.Printf("-----%v\n", reverseVowels("hello"))
}
