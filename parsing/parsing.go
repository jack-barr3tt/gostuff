package parsing

import (
	"fmt"
	"regexp"

	"github.com/jack-barr3tt/gostuff/types"
)

func NextTokenLocation(source string, exprs []regexp.Regexp) (*types.Pair[int, int], bool) {
	loc := types.Pair[int, int]{First: -1, Second: -1}

	for _, expr := range exprs {
		if match := expr.FindString(source); match != "" {
			if loc.First == -1 || expr.FindStringIndex(source)[0] < loc.First {
				locA := expr.FindStringIndex(source)
				loc = types.Pair[int, int]{First: locA[0], Second: locA[1]}
			}
		}
	}

	if loc.First == -1 {
		return nil, false
	}

	return &loc, true
}

func NextToken(source string, exprs []regexp.Regexp) (string, string) {
	loc, ok := NextTokenLocation(source, exprs)

	if !ok {
		return source, ""
	}

	return source[loc.Second:], source[loc.First:loc.Second]
}

func NextTokenStrict(source string, exprs []regexp.Regexp) (string, string, error) {
	loc, ok := NextTokenLocation(source, exprs)

	if !ok {
		return source, "", fmt.Errorf("no token found")
	}

	if loc.First != 0 {
		return source, "", fmt.Errorf("unexpected token")
	}

	return source[loc.Second:], source[:loc.Second], nil
}
