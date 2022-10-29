package properties

import (
	"fmt"
	"strings"

	"github.com/foxesknow/go-echo/collections"
)

const blockMarker = "\"\"\""

func FromString(config *Config, text string) (*Tree, error) {
	lines := strings.Split(text, "\n")

	root := newTree(config)
	activeSection := root

	for lineNo := 0; lineNo < len(lines); lineNo++ {
		line := strings.TrimSpace(lines[lineNo])

		// Skip anything that looks like whitespace
		if line == "" || isComment(config, line) {
			continue
		}

		// We'll allow properties at the top that are not part of a section to go into the root
		if sectionName, isSection := isSection(line); isSection {
			if sectionName == "" {
				return nil, fmt.Errorf("empty section name on line %d", lineNo+1)
			}

			sectionParts, err := getSectionParts(sectionName)
			if err != nil {
				return nil, err
			}

			// We need to navigate from the root down
			activeSection = root
			for _, sectionPart := range sectionParts {
				activeSection = activeSection.getOrCreateChild(sectionPart)
			}

		} else {
			key, value, lineDelta, err := extractKeyValue(lineNo, line, lines)
			if err != nil {
				return nil, err
			}

			// Apply any optional mappings that we've been configured for
			if keyMapper := root.config.KeyMapper; keyMapper != nil {
				key = keyMapper(key)
			}

			if valueMapper := root.config.ValueMapper; valueMapper != nil {
				value = valueMapper(value)
			}

			normalizedKey := root.config.caseNormalize(key)
			activeSection.values[normalizedKey] = collections.KeyValuePair[string, string]{Key: key, Value: value}

			lineNo += lineDelta
		}
	}

	return root, nil
}

func extractKeyValue(lineNo int, line string, lines []string) (key, value string, lineDelta int, err error) {
	pivot := strings.Index(line, "=")
	if pivot == -1 {
		return "", "", 0, fmt.Errorf("no = symbol on line %d", lineNo)
	}

	key = strings.TrimSpace(line[:pivot])
	value = strings.TrimSpace(line[pivot+1:])

	if value == blockMarker {
		value = ""
		lineDelta = 0

		foundMarker := false
		lineNo++
		for ; lineNo < len(lines); lineNo++ {
			s := lines[lineNo]
			if s == blockMarker {
				foundMarker = true
				break
			}

			lineDelta++
			if value == "" {
				value = s
			} else {
				value += "\n" + s
			}
		}

		if !foundMarker {
			return "", "", 0, fmt.Errorf("could not find end marker for %s", key)
		}
	}

	return key, value, lineDelta, nil
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

func getSectionParts(name string) ([]string, error) {
	parts := strings.Split(name, ".")

	for i, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, fmt.Errorf("invalid section name: %s", name)
		}

		parts[i] = part
	}

	return parts, nil
}
