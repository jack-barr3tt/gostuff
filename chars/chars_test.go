package chars

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestCharInRange(t *testing.T) {
	test.AssertEqual(t, CharInRange('a', 'z', 'a'), true)
	test.AssertEqual(t, CharInRange('b', 'z', 'a'), true)
	test.AssertEqual(t, CharInRange('A', 'a', 'z'), false)
	test.AssertEqual(t, CharInRange('0', 'a', 'z'), false)
}

func TestCharIsDigit(t *testing.T) {
	test.AssertEqual(t, CharIsDigit('0'), true)
	test.AssertEqual(t, CharIsDigit('9'), true)
	test.AssertEqual(t, CharIsDigit('a'), false)
	test.AssertEqual(t, CharIsDigit('A'), false)
}

func TestCharIsLower(t *testing.T) {
	test.AssertEqual(t, CharIsLower('a'), true)
	test.AssertEqual(t, CharIsLower('z'), true)
	test.AssertEqual(t, CharIsLower('A'), false)
	test.AssertEqual(t, CharIsLower('0'), false)
}

func TestCharIsUpper(t *testing.T) {
	test.AssertEqual(t, CharIsUpper('A'), true)
	test.AssertEqual(t, CharIsUpper('Z'), true)
	test.AssertEqual(t, CharIsUpper('a'), false)
	test.AssertEqual(t, CharIsUpper('0'), false)
}
