package collections

import (
	"sync"

	"github.com/foxesknow/go-echo/data"
)

type syncMap[K comparable, V any] struct {
	lock sync.RWMutex
	m    map[K]V
}

// Creates a new map that is thread safe
func NewSyncMap[K comparable, V any]() Map[K, V] {
	return &syncMap[K, V]{
		lock: sync.RWMutex{},
		m:    make(map[K]V),
	}
}

// Returns the number of items in the map
func (self *syncMap[K, V]) Count() int {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return len(self.m)
}

func (self *syncMap[K, V]) IsEmpty() bool {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return len(self.m) == 0
}

func (self *syncMap[K, V]) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.m = make(map[K]V)
}

// Checks to see if a key is present in the map
func (self *syncMap[K, V]) ContainsKey(key K) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()

	_, found := self.m[key]
	return found
}

// Tries to remove a key from the map.
// Return true if the key was removed, otherwise false.
func (self *syncMap[K, V]) Remove(key K) bool {
	self.lock.Lock()
	defer self.lock.Unlock()

	if _, found := self.m[key]; found {
		delete(self.m, key)
		return true
	}

	return false
}

// Attempts to get a value from the map.
// If the key does not exist returns (zero, false)
func (self *syncMap[K, V]) Get(key K) (item V, found bool) {
	self.lock.RLock()
	defer self.lock.RUnlock()

	item, found = self.m[key]
	return
}

// Adds an item to the map if the key does not already exists.
// Returns true if the item was added, false if not
func (self *syncMap[K, V]) Add(key K, value V) bool {
	self.lock.Lock()
	defer self.lock.Unlock()

	if _, found := self.m[key]; found {
		return false
	}

	self.m[key] = value
	return true
}

// Adds an item, if the key is not already present.
// If the key is present then the value is updated
func (self *syncMap[K, V]) AddOrUpdate(key K, value V) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.m[key] = value
}

// Returns a stream containing they keys in the map
func (self *syncMap[K, V]) Keys() data.Stream[K] {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return data.FromMapKeys(self.m)
}

// Returns a stream containing the values in the map
func (self *syncMap[K, V]) Values() data.Stream[V] {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return data.FromMapValues(self.m)
}

// Returns a stream containing they KeyValue pairs in the map
func (self *syncMap[K, V]) KeyValuePairs() data.Stream[data.KeyValuePair[K, V]] {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return data.FromMap(self.m)
}
