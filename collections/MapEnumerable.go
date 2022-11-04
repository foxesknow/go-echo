package collections

func extractKeys[K comparable, V any](m map[K]V) []K {
	slice := make([]K, 0, len(m))

	for k := range m {
		slice = append(slice, k)
	}

	return slice
}

func extractValues[K comparable, V any](m map[K]V) []V {
	slice := make([]V, 0, len(m))

	for _, v := range m {
		slice = append(slice, v)
	}

	return slice
}

func extractPairs[K comparable, V any](m map[K]V) []KeyValuePair[K, V] {
	slice := make([]KeyValuePair[K, V], len(m))

	for k, v := range m {
		slice = append(slice, KeyValuePair[K, V]{k, v})
	}

	return slice
}

func EnumerateMapKeys[K comparable, V any](m map[K]V) Enumerable[K] {
	return &FunctionEnumerable[K]{
		OnGetEnumerator: func() Enumerator[K] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractKeys(m)

			return &sliceEnumerator[K]{slice: slice, next: -1}
		},
	}
}

func EnumerateMapValues[K comparable, V any](m map[K]V) Enumerable[V] {
	return &FunctionEnumerable[V]{
		OnGetEnumerator: func() Enumerator[V] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractValues(m)

			return &sliceEnumerator[V]{slice: slice, next: -1}
		},
	}
}

func EnumerateKeyValuePairs[K comparable, V any](m map[K]V) Enumerable[KeyValuePair[K, V]] {
	return &FunctionEnumerable[KeyValuePair[K, V]]{
		OnGetEnumerator: func() Enumerator[KeyValuePair[K, V]] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractPairs(m)

			return &sliceEnumerator[KeyValuePair[K, V]]{slice: slice, next: -1}
		},
	}
}
