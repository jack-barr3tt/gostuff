package nums

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
