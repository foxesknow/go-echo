package settings

type SettingProvider interface {
	GetSetting(name string) (value string, found bool)
}
