package nums

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestGcd(t *testing.T) {
	test.AssertEqual(t, Gcd(0, 0), 0)

	test.AssertEqual(t, Gcd(6, 27), 3)

	test.AssertEqual(t, Gcd(27, 27), 27)
}

func TestLcm(t *testing.T) {
	test.AssertEqual(t, Lcm(6, 27), 54)

	test.AssertEqual(t, Lcm(55, 121), 605)
}

func TestFindLCM(t *testing.T) {
	test.AssertEqual(t, FindLCM([]int{}), 0)

	test.AssertEqual(t, FindLCM([]int{6, 27}), 54)

	test.AssertEqual(t, FindLCM([]int{6, 27, 55, 121}), 32670)
}
