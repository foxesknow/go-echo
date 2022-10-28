package ini

type Ini struct {
	sections map[string]*Section
	config   Config
}

func (self *Ini) Section(name string) (*Section, bool) {
	name = self.config.caseNormalize(name)
	section, found := self.sections[name]

	return section, found
}

func (self *Ini) Names() []string {
	names := make([]string, len(self.sections))

	i := 0
	for name := range self.sections {
		names[i] = name
		i++
	}

	return names
}
