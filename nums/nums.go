package nums

import "math"

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return (a * b) / Gcd(a, b)
}

func FindLCM(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = Lcm(result, numbers[i])
	}
	return result
}

func Abs[T ~int | ~float64 | ~int64 | ~float32](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Max[T ~int | ~float64 | ~int64 | ~float32](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Rationalize[T ~int | ~float64 | ~int64 | ~float32](n T, maxDenominator int) (int, int) {
	a, b := int(n), 1
	diff := float64(1 << 31)

	for i := 1; i <= maxDenominator; i++ {
		newA := int(math.Round(float64(n) * float64(i)))
		if newDiff := Abs(float64(n) - float64(newA)/float64(i)); newDiff < diff {
			a, b, diff = newA, i, newDiff
		}
		if diff == 0 {
			break
		}
	}

	return a, b
}

func Pow(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}
