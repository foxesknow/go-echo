package settings

import (
	"fmt"
	"strings"

	"github.com/foxesknow/go-echo/settings/internal/providers"
)

var allSettings = make(map[string]SettingProvider)

func init() {
	allSettings["env"] = providers.NewEnvironmentSettings()
	allSettings["os"] = providers.NewOSSettings()
}

func IsRegistered(provider string) bool {
	_, present := allSettings[normalizeName(provider)]
	return present
}

// Returns the value for a setting.
// The setting must be in Provider:Name format
func Value(setting string) (string, bool) {
	provider, name, err := extractProviderAndName(setting)
	if err != nil {
		return "", false
	}

	p, ok := allSettings[normalizeName(provider)]
	if !ok {
		return "", false
	}

	return p.GetSetting(name)
}

func normalizeName(name string) string {
	return strings.ToLower(name)
}

func extractProviderAndName(value string) (string, string, error) {
	pivot := strings.Index(value, ":")

	if pivot == -1 {
		return "", "", fmt.Errorf("could not find :")
	}

	provider := value[:pivot]
	setting := value[pivot+1:]

	return provider, setting, nil
}
