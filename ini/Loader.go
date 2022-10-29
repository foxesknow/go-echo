package ini

import (
	"bytes"
	"fmt"
	"go-echo/collections"
	"io"
	"os"
	"strings"
)

// Loads an ini file from a file
func FromFile(config *Config, name string) (*Ini, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("could not open %s", name)
	}

	defer file.Close()

	ini, err := FromReader(config, file)
	if err != nil {
		return nil, fmt.Errorf("%w; error processing %s", err, name)
	}

	return ini, nil
}

// Loads an ini file from a reader
func FromReader(config *Config, reader io.Reader) (*Ini, error) {
	buffer := bytes.Buffer{}

	_, err := buffer.ReadFrom(reader)
	if err != nil {
		return nil, err
	}

	text := buffer.String()
	return FromString(config, text)
}

// Loads an ini file from test
func FromString(config *Config, text string) (*Ini, error) {
	lines := strings.Split(text, "\n")

	ini := newIni(config)

	var activeSection *Section

	for lineNo := 0; lineNo < len(lines); lineNo++ {
		line := strings.TrimSpace(lines[lineNo])

		// Skip anything that looks like whitespace
		if line == "" || isComment(config, line) {
			continue
		}

		if sectionName, isSection := isSection(line); isSection {
			if sectionName == "" {
				return nil, fmt.Errorf("empty section name on line %d", lineNo+1)
			}

			// We'll keep the section name in the case it arrived in for the name of the section
			// but store it using the casing rules
			normalizedSectionName := ini.config.caseNormalize(sectionName)
			activeSection = newSection(config, sectionName)
			ini.sections[normalizedSectionName] = activeSection

		} else {
			// It's a key-value pair
			if activeSection == nil {
				return nil, fmt.Errorf("found a key when not is a section on line %d", lineNo+1)
			}

			key, value, err := extractKeyValue(lineNo, line)
			if err != nil {
				return nil, err
			}

			// Apply any optional mappings that we've been configured for
			if keyMapper := ini.config.KeyMapper; keyMapper != nil {
				key = keyMapper(key)
			}

			if valueMapper := ini.config.ValueMapper; valueMapper != nil {
				value = valueMapper(value)
			}

			normalizedKey := ini.config.caseNormalize(key)
			activeSection.values[normalizedKey] = collections.KeyValuePair[string, string]{Key: key, Value: value}
		}
	}

	return ini, nil
}

func extractKeyValue(lineNo int, line string) (key, value string, err error) {
	pivot := strings.Index(line, "=")
	if pivot == -1 {
		return "", "", fmt.Errorf("no = symbol on line %d", lineNo)
	}

	key = strings.TrimSpace(line[:pivot])
	value = strings.TrimSpace(line[pivot+1:])
	err = nil

	return
}

func isComment(config *Config, text string) bool {
	return text[0] == ';' || text[0] == '#'
}

func isSection(text string) (string, bool) {
	if text[0] == '[' && text[len(text)-1] == ']' {
		// We'll let the caller decide what to do with an empty section name
		name := text[1 : len(text)-1]
		return strings.TrimSpace(name), true
	}

	return "", false
}

func newIni(config *Config) *Ini {
	return &Ini{
		sections: make(map[string]*Section),
		config:   *config,
	}
}
