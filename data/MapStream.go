package data

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
		slice = append(slice, KeyValuePair[K, V]{Key: k, Value: v})
	}

	return slice
}

// Returns a stream for the keys in a map
func FromMapKeys[K comparable, V any](m map[K]V) Stream[K] {
	return &FunctionStream[K]{
		OnIterator: func() Iterator[K] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractKeys(m)

			return &sliceIterator[K]{slice: slice, next: -1}
		},
	}
}

// Returns a stream for the values in a map
func FromMapValues[K comparable, V any](m map[K]V) Stream[V] {
	return &FunctionStream[V]{
		OnIterator: func() Iterator[V] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractValues(m)

			return &sliceIterator[V]{slice: slice, next: -1}
		},
	}
}

// Returns a stream for the key/value pairs in a map
func FromMap[K comparable, V any](m map[K]V) Stream[KeyValuePair[K, V]] {
	return &FunctionStream[KeyValuePair[K, V]]{
		OnIterator: func() Iterator[KeyValuePair[K, V]] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractPairs(m)

			return &sliceIterator[KeyValuePair[K, V]]{slice: slice, next: -1}
		},
	}
}
