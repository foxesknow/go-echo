package collections

import (
	"testing"
)

func Test_SyncMap_Empty(t *testing.T) {
	m := NewSyncMap[int, string]()

	if m.Count() != 0 {
		t.Error("map should have 0 count")
	}

	if !m.IsEmpty() {
		t.Error("map should be empty")
	}
}

func Test_SyncMap_Count(t *testing.T) {
	m := NewSyncMap[int, string]()

	if m.Count() != 0 {
		t.Error("map should have 0 count")
	}

	m.Add(1, "Jack")

	if m.Count() != 1 {
		t.Error("map should have 1 count")
	}
}

func Test_SyncMap_IsEmpty(t *testing.T) {
	m := NewSyncMap[int, string]()

	if !m.IsEmpty() {
		t.Error("map should be empty")
	}

	m.Add(1, "Jack")

	if m.IsEmpty() {
		t.Error("map should not be empty")
	}

	m.Clear()

	if !m.IsEmpty() {
		t.Error("map should be empty now")
	}
}

func Test_SyncMap_ContainsKey(t *testing.T) {
	m := NewSyncMap[int, string]()

	if !m.IsEmpty() {
		t.Error("map should be empty")
	}

	if m.ContainsKey(1) {
		t.Error("1 should not be present")
	}

	m.Add(1, "Jack")

	if !m.ContainsKey(1) {
		t.Error("1 should be present")
	}
}

func Test_SyncMap_Remove(t *testing.T) {
	m := NewSyncMap[int, string]()

	if !m.IsEmpty() {
		t.Error("map should be empty")
	}

	if m.Remove(1) {
		t.Error("1 should not be present")
	}

	m.Add(1, "Jack")

	if !m.Remove(1) {
		t.Error("should have removed 1")
	}
}

func Test_SyncMap_Get(t *testing.T) {
	m := NewSyncMap[int, string]()

	if _, found := m.Get(1); found {
		t.Error("1 is not in the map")
	}

	m.Add(1, "Jack")

	if item, found := m.Get(1); !found || item != "Jack" {
		t.Error("unexpected return")
	}
}

func Test_SyncMap_Add(t *testing.T) {
	m := NewSyncMap[int, string]()

	if !m.Add(1, "Jack") {
		t.Error("we should have been able to add map")
	}

	if m.Add(1, "Jack") {
		t.Error("1 is already mapped!")
	}

	if item, _ := m.Get(1); item != "Jack" {
		t.Error("1 has changed!!")
	}
}

func Test_SyncMap_AddOrUpdate(t *testing.T) {
	m := NewSyncMap[int, string]()

	m.AddOrUpdate(1, "Jack")

	if item, _ := m.Get(1); item != "Jack" {
		t.Error("1 should map to Jack")
	}

	m.AddOrUpdate(1, "Sawyer")

	if item, _ := m.Get(1); item != "Sawyer" {
		t.Error("1 should map to Sawyer")
	}
}

func Test_SyncMap_Mix(t *testing.T) {
	m := NewSyncMap[int, string]()

	m.Add(1, "Jack")
	m.AddOrUpdate(2, "Sawyer")
	m.Add(3, "Ben")
	m.AddOrUpdate(4, "Hurley")

	if m.Count() != 4 {
		t.Error("should have 4 items")
	}

	for i := 1; i <= 4; i++ {
		if _, found := m.Get(1); !found {
			t.Errorf("could not find %d", i)
		}
	}
}

func Test_SyncMap_Keys(t *testing.T) {
	m := NewSyncMap[int, string]()

	m.Add(1, "Jack")
	m.AddOrUpdate(2, "Sawyer")
	m.Add(3, "Ben")
	m.AddOrUpdate(4, "Hurley")

	it := m.Keys().GetStream()

	for i := 1; i <= 4; i++ {
		if !it.MoveNext() {
			t.Error("should have been able to move")
		}
	}

	// There shouldn't be anything left
	if it.MoveNext() {
		t.Error("we should be at the end of the iterator")
	}
}

func Test_SyncMap_Values(t *testing.T) {
	m := NewSyncMap[int, string]()

	m.Add(1, "Jack")
	m.AddOrUpdate(2, "Sawyer")
	m.Add(3, "Ben")
	m.AddOrUpdate(4, "Hurley")

	it := m.Values().GetStream()

	for i := 1; i <= 4; i++ {
		if !it.MoveNext() {
			t.Error("should have been able to move")
		}
	}

	// There shouldn't be anything left
	if it.MoveNext() {
		t.Error("we should be at the end of the iterator")
	}
}

func Test_SyncMap_KeyValuePairs(t *testing.T) {
	m := NewSyncMap[int, string]()

	m.Add(1, "Jack")
	m.AddOrUpdate(2, "Sawyer")
	m.Add(3, "Ben")
	m.AddOrUpdate(4, "Hurley")

	it := m.KeyValuePairs().GetStream()

	for i := 1; i <= 4; i++ {
		if !it.MoveNext() {
			t.Error("should have been able to move")
		}
	}

	// There shouldn't be anything left
	if it.MoveNext() {
		t.Error("we should be at the end of the iterator")
	}
}
