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
