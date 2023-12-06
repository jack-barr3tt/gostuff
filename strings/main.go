package strings

import (
	"regexp"
	"strconv"

	"github.com/jack-barr3tt/gostuff/slices"
)

var numregex = regexp.MustCompile(`\d+`)

func GetNum(s string) int {
	num, _ := strconv.Atoi(numregex.FindString(s))
	return num
}

func GetNums(s string) []int {
	numstrs := numregex.FindAllString(s, -1)

	return slices.StrsToInts(numstrs)
}
