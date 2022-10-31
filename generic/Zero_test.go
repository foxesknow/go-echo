package generic

import (
	"testing"
)

func TestZero(t *testing.T) {
	if value := Zero[bool](); value != false {
		t.Error("expected false")
	}

	if value := Zero[int8](); value != 0 {
		t.Error("expected 0")
	}

	if value := Zero[int16](); value != 0 {
		t.Error("expected 0")
	}

	if value := Zero[int32](); value != 0 {
		t.Error("expected 0")
	}

	if value := Zero[int64](); value != 0 {
		t.Error("expected 0")
	}

	if value := Zero[int](); value != 0 {
		t.Error("expected 0")
	}

	if value := Zero[float32](); value != 0.0 {
		t.Error("expected 0")
	}

	if value := Zero[float64](); value != 0.0 {
		t.Error("expected 0")
	}

	if value := Zero[map[string]string](); value != nil {
		t.Error("expected nil")
	}

	if value := Zero[[]int](); value != nil {
		t.Error("expected nil")
	}
}
