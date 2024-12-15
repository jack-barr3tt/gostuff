package types

type Pair[A, B comparable] struct {
	First  A
	Second B
}

func (p Pair[A, B]) Equals(other Pair[A, B]) bool {
	return p.First == other.First && p.Second == other.Second
}

func PairFromSlice[A comparable](s []A) Pair[A, A] {
	return Pair[A, A]{First: s[0], Second: s[1]}
}
