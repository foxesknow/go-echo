package environment

import "os"

type EnvironmentSettings struct {
}

func New() *EnvironmentSettings {
	return &EnvironmentSettings{}
}

func (self *EnvironmentSettings) GetSetting(name string) (value string, found bool) {
	value, found = os.LookupEnv(name)
	return
}
