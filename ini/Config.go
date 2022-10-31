package ini

import "strings"

// Configuration information for an ini file
type Config struct {
	// Whether section names and keys should be case sensitive
	CaseSensitive bool

	// (optional) called on each key during the loading process
	KeyMapper func(key string) string

	// (optional) called on each value during the loading process
	// If a KeyMapper is present then "key" holds the mapped value
	ValueMapper func(key string, value string) string
}

// Applies the casing rules as specified by the user
func (self *Config) caseNormalize(text string) string {
	if !self.CaseSensitive {
		return strings.ToLower(text)
	}

	return text
}
