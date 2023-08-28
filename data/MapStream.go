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
	slice := make([]KeyValuePair[K, V], 0, len(m))

	for k, v := range m {
		slice = append(slice, KeyValuePair[K, V]{Key: k, Value: v})
	}

	return slice
}

// Returns a stream for the keys in a map
func FromMapKeys[K comparable, V any](m map[K]V) Streamable[K] {
	return &FunctionStreamable[K]{
		OnGetStream: func() Stream[K] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractKeys(m)

			return &sliceStream[K]{slice: slice, next: -1}
		},
	}
}

// Returns a stream for the values in a map
func FromMapValues[K comparable, V any](m map[K]V) Streamable[V] {
	return &FunctionStreamable[V]{
		OnGetStream: func() Stream[V] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractValues(m)

			return &sliceStream[V]{slice: slice, next: -1}
		},
	}
}

// Returns a stream for the key/value pairs in a map
func FromMap[K comparable, V any](m map[K]V) Streamable[KeyValuePair[K, V]] {
	return &FunctionStreamable[KeyValuePair[K, V]]{
		OnGetStream: func() Stream[KeyValuePair[K, V]] {
			// Extract the keys here so we only do so if the user does enumerate
			slice := extractPairs(m)

			return &sliceStream[KeyValuePair[K, V]]{slice: slice, next: -1}
		},
	}
}
