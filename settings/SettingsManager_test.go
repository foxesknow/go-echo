package settings

import (
	"testing"
)

func TestIsRegistered(t *testing.T) {
	registered := IsRegistered("env")
	if !registered {
		t.Error("env not registered")
	}
}

func TestScopedName(t *testing.T) {
	value, found := Value("os:homedir")
	if !found {
		t.Error("could not resolve setting")
		return
	}

	if value == "" {
		t.Error("value is empty")
		return
	}
}
