package settings

type ExplicitSettings struct {
	items map[string]string
}

func NewExplicitSettings() *ExplicitSettings {
	return &ExplicitSettings{items: make(map[string]string)}
}

func (self *ExplicitSettings) Add(name, value string) {
	self.items[name] = value
}

func (self *ExplicitSettings) IsRegistered(name string) bool {
	_, found := self.items[name]
	return found
}
