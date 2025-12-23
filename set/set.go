package set

type Set[T comparable] struct {
	data map[T]bool
}

func FromSlice[T comparable](items []T) Set[T] {
	s := Set[T]{data: make(map[T]bool)}
	for _, item := range items {
		s.data[item] = true
	}
	return s
}

func FromMap[K comparable, V any](m map[K]V) Set[K] {
	s := Set[K]{data: make(map[K]bool)}
	for k := range m {
		s.data[k] = true
	}
	return s
}

func (s Set[T]) ToSlice() []T {
	items := make([]T, 0, len(s.data))
	for item := range s.data {
		items = append(items, item)
	}
	return items
}

func (s Set[T]) Add(item T) {
	s.data[item] = true
}

func (s Set[T]) Remove(item T) {
	delete(s.data, item)
}

func (s Set[T]) Has(item T) bool {
	_, exists := s.data[item]
	return exists
}

func (s Set[T]) Size() int {
	return len(s.data)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := FromSlice(s.ToSlice())
	for item := range other.data {
		result.data[item] = true
	}
	return result
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := Set[T]{data: make(map[T]bool)}
	for item := range s.data {
		if other.Has(item) {
			result.data[item] = true
		}
	}
	return result
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := Set[T]{data: make(map[T]bool)}
	for item := range s.data {
		if !other.Has(item) {
			result.data[item] = true
		}
	}
	return result
}

func (s Set[T]) IsSubset(other Set[T]) bool {
	for item := range s.data {
		if !other.Has(item) {
			return false
		}
	}
	return true
}
