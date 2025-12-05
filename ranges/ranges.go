package ranges

import (
	"cmp"
	"regexp"
	"slices"

	"github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func CombineRanges(ranges []types.Range) []types.Range {
	slices.SortFunc(ranges, func(a, b types.Range) int {
		return cmp.Compare(a.Start, b.Start)
	})

	combined := []types.Range{}

	current := ranges[0]

	for _, r := range ranges[1:] {
		if current.OverlapsRange(r) {
			current = *current.AddRange(r)
		} else {
			combined = append(combined, current)
			current = r
		}
	}

	combined = append(combined, current)

	return combined
}

var rangeRegex = regexp.MustCompile(`^\s*(-?\d+)\s*-\s*(-?\d+)\s*$`)

func ParseRange(s string) types.Range {
	matches := rangeRegex.FindStringSubmatch(s)
	if matches == nil {
		panic("Invalid range string: " + s)
	}

	return types.Range{Start: strings.GetNum(matches[1]), End: strings.GetNum(matches[2])}
}
