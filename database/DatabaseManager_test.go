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

func Test_Register_NoAlias(t *testing.T) {
	manager := NewDatabaseManager()

	info := ConnectionInfo{DriverName: "testdb", ConnectionString: "something"}
	err := manager.Register("", info)

	if err == nil {
		t.Errorf("Register should fail")
	}
}

func Test_Register_NoDriver(t *testing.T) {
	manager := NewDatabaseManager()

	info := ConnectionInfo{ConnectionString: "something"}
	err := manager.Register("foo", info)

	if err == nil {
		t.Errorf("Register should fail")
	}
}

func Test_Register_NoConnectionString(t *testing.T) {
	manager := NewDatabaseManager()

	info := ConnectionInfo{DriverName: "testdb"}
	err := manager.Register("foo", info)

	if err == nil {
		t.Errorf("Register should fail")
	}
}

func Test_Register_Already(t *testing.T) {
	manager := NewDatabaseManager()

	info := ConnectionInfo{DriverName: "testdb", ConnectionString: "something"}

	if err := manager.Register("foo", info); err != nil {
		t.Errorf("Register should succeed")
	}

	if err := manager.Register("foo", info); err == nil {
		t.Errorf("Register should fail")
	}
}
