package collections

import "github.com/foxesknow/go-echo/data"

type defaultMap[K comparable, V any] struct {
	m map[K]V
}

// Creates a new map that is not thread safe
func NewMap[K comparable, V any]() Map[K, V] {
	return &defaultMap[K, V]{
		m: make(map[K]V),
	}
}

// Returns the number of items in the map
func (self *defaultMap[K, V]) Count() int {
	return len(self.m)
}

func (self *defaultMap[K, V]) IsEmpty() bool {
	return len(self.m) == 0
}

func (self *defaultMap[K, V]) Clear() {
	self.m = make(map[K]V)
}

// Checks to see if a key is present in the map
func (self *defaultMap[K, V]) ContainsKey(key K) bool {
	_, found := self.m[key]
	return found
}

// Tries to remove a key from the map.
// Return true if the key was removed, otherwise false.
func (self *defaultMap[K, V]) Remove(key K) bool {
	if _, found := self.m[key]; found {
		delete(self.m, key)
		return true
	}

	return false
}

// Attempts to get a value from the map.
// If the key does not exist returns (zero, false)
func (self *defaultMap[K, V]) Get(key K) (item V, found bool) {
	item, found = self.m[key]
	return
}

// Adds an item to the map if the key does not already exists.
// Returns true if the item was added, false if not
func (self *defaultMap[K, V]) Add(key K, value V) bool {
	if self.ContainsKey(key) {
		return false
	}

	self.m[key] = value
	return true
}

// Adds an item, if the key is not already present.
// If the key is present then the value is updated
func (self *defaultMap[K, V]) AddOrUpdate(key K, value V) {
	self.m[key] = value
}

// Returns a stream containing they keys in the map
func (self *defaultMap[K, V]) Keys() data.Stream[K] {
	return data.FromMapKeys(self.m)
}

// Returns a stream containing the values in the map
func (self *defaultMap[K, V]) Values() data.Stream[V] {
	return data.FromMapValues(self.m)
}

// Returns a stream containing they KeyValue pairs in the map
func (self *defaultMap[K, V]) KeyValuePairs() data.Stream[data.KeyValuePair[K, V]] {
	return data.FromMap(self.m)
}
