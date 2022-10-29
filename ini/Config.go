package ini

import "strings"

// Configiration information for an ini file
type Config struct {
	// Whether section names and keys should be case sensitive
	CaseSensitive bool

	// (optional) called on each key during the loading process
	KeyMapper KeyMapper

	// (optional) called on each value during the loading process
	ValueMapper ValueMapper
}

// Applies the casing rules as specified by the user
func (self *Config) caseNormalize(text string) string {
	if !self.CaseSensitive {
		return strings.ToLower(text)
	}

	return text
}
