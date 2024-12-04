package parsing

import (
	"regexp"
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
	"github.com/jack-barr3tt/gostuff/types"
)

var print = regexp.MustCompile(`print\('.+?'\)`)
var okr = regexp.MustCompile(`ok\(\d+\)`)
var add = regexp.MustCompile(`\d+\+\d+`)

var testSrc = "'#;print('hi');';l23;ok(1)kjshj"

func TestNextTokenLocation(t *testing.T) {
	// test match
	loc, ok := NextTokenLocation(testSrc, []regexp.Regexp{*print, *okr})
	test.AssertEqual(t, ok, true)
	test.AssertEqual(t, loc, &types.Pair[int, int]{First: 3, Second: 14})

	// test different match
	loc, ok = NextTokenLocation(testSrc, []regexp.Regexp{*okr})
	test.AssertEqual(t, ok, true)
	test.AssertEqual(t, loc, &types.Pair[int, int]{First: 21, Second: 26})

	// test no match
	loc, ok = NextTokenLocation(testSrc, []regexp.Regexp{*add})
	test.AssertEqual(t, ok, false)
	test.AssertEqual(t, loc, nil)
}

func TestNextToken(t *testing.T) {
	exprs := []regexp.Regexp{*print, *okr}

	// test match
	rest, token := NextToken(testSrc, exprs)
	test.AssertEqual(t, token, "print('hi')")
	test.AssertEqual(t, rest, ";';l23;ok(1)kjshj")

	// test another match
	rest, token = NextToken(rest, exprs)
	test.AssertEqual(t, token, "ok(1)")
	test.AssertEqual(t, rest, "kjshj")

	// test no match
	rest, token = NextToken(rest, exprs)
	test.AssertEqual(t, token, "")
	test.AssertEqual(t, rest, "kjshj")
}

func TestNextTokenStrict(t *testing.T) {
	exprs := []regexp.Regexp{*print, *okr}

	// test match
	rest, token, err := NextTokenStrict("print('hi')ok(1)kjshj", exprs)
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, token, "print('hi')")
	test.AssertEqual(t, rest, "ok(1)kjshj")

	// test another match
	rest, token, err = NextTokenStrict(rest, exprs)
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, token, "ok(1)")
	test.AssertEqual(t, rest, "kjshj")

	// test no match
	rest, token, err = NextTokenStrict(rest, exprs)
	test.AssertEqual(t, err.Error(), "no token found")
	test.AssertEqual(t, token, "")
	test.AssertEqual(t, rest, "kjshj")
}
