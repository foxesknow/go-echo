package ini

import "strings"

type Section struct {
	name   string
	values map[string]string
	config Config
}

func newSection(config *Config, name string) *Section {
	return &Section{
		name:   name,
		values: make(map[string]string),
		config: *config,
	}
}

// Returns the number of items in the section
func (self *Section) Len() int {
	return len(self.values)
}

// Returns the name of the section, in the case it existed in the source
func (self *Section) Name() string {
	return self.name
}

// Returns the value for a given key.
// If the key does not exist returns ("", false)
func (self *Section) Value(key string) (value string, found bool) {
	key = self.normalizeKey(key)
	value, found = self.values[key]
	return
}

// Returns the value for a given key, or defaultValue if it does not exist
func (self *Section) ValueOrDefault(key string, defaultValue string) string {
	key = self.normalizeKey(key)
	value, found := self.values[key]

	if found {
		return value
	} else {
		return defaultValue
	}
}

// Checks to see if a key is present
func (self *Section) HasKey(key string) bool {
	key = self.normalizeKey(key)
	_, found := self.values[key]
	return found
}

// Returns the name of all the keys in the section
func (self *Section) Keys() []string {
	keys := make([]string, len(self.values))

	var i = 0
	for key := range self.values {
		keys[i] = key
		i++
	}

	return keys
}

// Returns the key/value pairs for all items in the section
func (self *Section) Pairs() []KeyValue {
	pairs := make([]KeyValue, len(self.values))

	var i = 0
	for key, value := range self.values {
		pairs[i] = KeyValue{key, value}
		i++
	}

	return pairs
}

func (self *Section) normalizeKey(key string) string {
	if !self.config.CaseSensitive {
		return strings.ToLower(key)
	}

	return key
}

// Implement Stringer
func (self *Section) String() string {
	return self.name
}
