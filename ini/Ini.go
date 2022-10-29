package ini

type Ini struct {
	sections map[string]*Section
	config   Config
}

// Returns the specified section.
// If not found then returns (nil, false)
func (self *Ini) Section(name string) (section *Section, found bool) {
	name = self.config.caseNormalize(name)
	section, found = self.sections[name]
	return
}

// Returns the names of all the sections in the ini file
func (self *Ini) Names() []string {
	names := make([]string, len(self.sections))

	i := 0
	for name := range self.sections {
		names[i] = name
		i++
	}

	return names
}
