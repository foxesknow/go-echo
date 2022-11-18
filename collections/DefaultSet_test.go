package collections

import (
	"testing"
)

func Test_DefaultSet_Empty(t *testing.T) {
	s := NewSet[string]()

	if s.Count() != 0 {
		t.Error("set should have 0 count")
	}

	if !s.IsEmpty() {
		t.Error("set should be empty")
	}
}

func Test_DefaultSet_Add(t *testing.T) {
	s := NewSet[string]()

	if !s.Add("Ben") {
		t.Error("should have added Ben")
	}

	if !s.Add("Jack") {
		t.Error("should have added Jack")
	}

	if s.Count() != 2 {
		t.Error("set should have 2 items")
	}
}

func Test_DefaultSet_Add_Duplicate(t *testing.T) {
	s := NewSet[string]()

	if !s.Add("Ben") {
		t.Error("should have added Ben")
	}

	if !s.Add("Jack") {
		t.Error("should have added Jack")
	}

	if s.Add("Ben") {
		t.Error("should not have added Ben")
	}

	if s.Count() != 2 {
		t.Error("set should have 2 items")
	}
}

func Test_DefaultSet_Contains(t *testing.T) {
	s := NewSet[string]()

	if s.Contains("Jack") {
		t.Error("Jack shouldn't exist")
	}

	s.Add("Jack")

	if !s.Contains("Jack") {
		t.Error("Jack should exist")
	}
}

func Test_DefaultSet_Remove(t *testing.T) {
	s := NewSet[string]()

	s.Add("Ben")
	s.Add("Jack")

	if !s.Remove("Jack") {
		t.Error("should have removed Jack")
	}

	if s.Count() != 1 {
		t.Error("set should have 1 items")
	}

	// We should be able to add it back
	if !s.Add("Jack") {
		t.Error("Should have added Jack")
	}
}

func Test_DefaultSet_Remove_NotPresent(t *testing.T) {
	s := NewSet[string]()

	s.Add("Ben")
	s.Add("Jack")

	if s.Remove("Hurley") {
		t.Error("should not have removed Hurley")
	}
}

func Test_DefaultSet_Clear(t *testing.T) {
	s := NewSet[string]()

	s.Add("Ben")
	s.Add("Jack")

	if s.Count() != 2 {
		t.Error("should have 2 items")
	}

	s.Clear()

	if s.Count() != 0 {
		t.Error("should have 0 items")
	}
}

func Test_DefaultSet_Stream_Empty(t *testing.T) {
	s := NewSet[string]()

	stream := s.Stream()
	i := stream.Iterator()

	if i.MoveNext() {
		t.Error("there should be nothing in the set")
	}
}

func Test_DefaultSet_Stream(t *testing.T) {
	s := NewSet[string]()
	s.Add("Kate")

	stream := s.Stream()
	i := stream.Iterator()

	if !i.MoveNext() {
		t.Error("there should be something in the set")
	}

	if i.Current() != "Kate" {
		t.Error("expected Kate")
	}
}

func Test_DefaultSet_Union(t *testing.T) {
	lhs := NewSet[int]()
	lhs.Add(1)
	lhs.Add(5)
	lhs.Add(9)

	rhs := NewSet[int]()
	rhs.Add(2)
	rhs.Add(3)
	rhs.Add(4)

	lhs.Union(rhs.Stream())

	if lhs.Count() != 6 {
		t.Error("should have 6 items")
	}

	for _, value := range []int{1, 5, 9, 2, 3, 4} {
		if !lhs.Contains(value) {
			t.Errorf("expected %d", value)
		}
	}
}

func Test_DefaultSet_Union_Duplicates(t *testing.T) {
	lhs := NewSet[int]()
	lhs.Add(1)
	lhs.Add(5)
	lhs.Add(9)

	rhs := NewSet[int]()
	rhs.Add(2)
	rhs.Add(3)
	rhs.Add(9)

	lhs.Union(rhs.Stream())

	if lhs.Count() != 5 {
		t.Error("should have 5 items")
	}

	for _, value := range []int{1, 5, 9, 2, 3} {
		if !lhs.Contains(value) {
			t.Errorf("expected %d", value)
		}
	}
}

func Test_DefaultSet_Except(t *testing.T) {
	lhs := NewSet[int]()
	lhs.Add(1)
	lhs.Add(5)
	lhs.Add(9)

	rhs := NewSet[int]()
	rhs.Add(2)
	rhs.Add(3)
	rhs.Add(4)

	lhs.Except(rhs.Stream())

	if lhs.Count() != 3 {
		t.Error("should have 3 items")
	}

	for _, value := range []int{1, 5, 9} {
		if !lhs.Contains(value) {
			t.Errorf("expected %d", value)
		}
	}
}

func Test_DefaultSet_Excep_Matches(t *testing.T) {
	lhs := NewSet[int]()
	lhs.Add(1)
	lhs.Add(5)
	lhs.Add(9)

	rhs := NewSet[int]()
	rhs.Add(2)
	rhs.Add(5)
	rhs.Add(9)

	lhs.Except(rhs.Stream())

	if lhs.Count() != 1 {
		t.Error("should have 3 items")
	}

	for _, value := range []int{1} {
		if !lhs.Contains(value) {
			t.Errorf("expected %d", value)
		}
	}
}
