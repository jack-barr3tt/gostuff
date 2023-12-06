package main

import (
	"fmt"

	"github.com/jack-barr3tt/gostuff/chars"
	"github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

// I use this to test the packages 👍
func main() {
	raw := "1 2 3 4 5 6 7 8 9 10"

	nums := strings.GetNums(raw)

	fmt.Println("nums:", nums)

	evens := slices.Filter(func(num int) bool {
		return num%2 == 0
	}, nums)

	fmt.Println("evens:", evens)

	evensSum := slices.Reduce(func(num int, sum int) int {
		return sum + num
	}, evens, 0)

	fmt.Println("sum of evens:", evensSum)

	println(chars.CharIsDigit('1'))

	digits := types.Range{Start: 0, End: 9}

	fmt.Println("digits:", digits)

	println(digits.Contains(5))
	println(digits.Contains(10))

	fmt.Println(digits.SubtractRange(types.Range{Start: 3, End: 5}))
	fmt.Println(digits.SubtractRange(types.Range{Start: 0, End: 9}))
	fmt.Println(digits.SubtractRange(types.Range{Start: 0, End: 5}))
	fmt.Println(digits.SubtractRange(types.Range{Start: 3, End: 20}))

	fmt.Println(digits.SplitAfter(5))

	fmt.Println(slices.Zip([]int{1, 2, 3}, []string{"a", "b", "c"}))

	nums1 := []int{5, 9, 2, 4, 7, 1, 3, 8, 6}
	nums2 := []int{3, 8, 6, 5, 9, 2, 4, 7, 1}

	numsCombined := slices.Zip(nums1, nums2)

	numsIsGreater := slices.Map(func(nums types.Pair[int, int]) bool {
		return nums.First > nums.Second
	}, numsCombined)

	fmt.Println(numsIsGreater)

	multiNums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(slices.Flat(multiNums))

	fmt.Println(slices.Combos([]int{1, 2, 3}, []string{"a", "b", "c"}))
}
