package ini

import "strings"

type Config struct {
	CaseSensitive bool
	KeyMapper     KeyMapper
	ValueMapper   ValueMapper
}

func (self *Config) caseNormalize(text string) string {
	if !self.CaseSensitive {
		return strings.ToLower(text)
	}

	return text
}
