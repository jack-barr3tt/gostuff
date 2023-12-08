package maps

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
