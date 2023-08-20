package database

import (
	"testing"
)

func Test_Initialize(t *testing.T) {
	manager := NewDatabaseManager()

	db, err := manager.Open("foo")

	if err == nil {
		t.Errorf("nothing should exist")
	}

	if db != nil {
		t.Errorf("no database should exist")
	}
}
