package settings

import "go-echo/settings/internal/environment"

var providers = make(map[string]SettingProvider)

func init() {
	providers["env"] = environment.New()
}

func IsRegistered(provider string) bool {
	_, present := providers[provider]
	return present
}
