package maps

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0)
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Clone[K comparable, V any](m map[K]V) map[K]V {
	clone := make(map[K]V)
	for k, v := range m {
		clone[k] = v
	}
	return clone
}

func Map[K comparable, MV any, RV any](f func(K, MV) RV, m map[K]MV) []RV {
	result := []RV{}
	for k, v := range m {
		result = append(result, f(k, v))
	}
	return result
}

func Reduce[K comparable, V any, Acc any](f func(Acc, K, V) Acc, m map[K]V, initial Acc) Acc {
	result := initial
	for k, v := range m {
		result = f(result, k, v)
	}
	return result
}
