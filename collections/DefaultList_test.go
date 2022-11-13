package collections

import (
	"testing"
)

func Test_List_Empty(t *testing.T) {
	list := NewList[int]()

	if list.Count() != 0 {
		t.Error("list should have count of 0")
	}

	if !list.IsEmpty() {
		t.Error("list should be empty")
	}
}

func Test_Clear(t *testing.T) {
	list := NewList[int]()

	list.Add(10)
	list.Add(20)

	if list.Count() != 2 {
		t.Error("list should have count of 2")
	}

	list.Clear()

	if !list.IsEmpty() {
		t.Error("list isn't empty")
	}
}

func Test_Add(t *testing.T) {
	list := NewList[int]()

	list.Add(10)
	list.Add(20)

	if list.Count() != 2 {
		t.Error("list should have count of 2")
	}

	if list.IsEmpty() {
		t.Error("list isn't empty")
	}
}

func Test_Get_Empty(t *testing.T) {
	list := NewList[int]()

	_, err := list.Get(0)
	if err == nil {
		t.Error("Get should have returned an error")
	}
}

func Test_Get(t *testing.T) {
	list := NewList[int]()

	list.Add(10)
	list.Add(20)

	if item, err := list.Get(0); err != nil || item != 10 {
		t.Error("Expected 10")
	}

	if item, err := list.Get(1); err != nil || item != 20 {
		t.Error("Expected 20")
	}

	if _, err := list.Get(2); err == nil {
		t.Error("Shold have got an error")
	}
}

func Test_Set_Empty(t *testing.T) {
	list := NewList[int]()

	if err := list.Set(0, 99); err == nil {
		t.Error("Set should have returned an error")
	}
}

func Test_Set(t *testing.T) {
	list := NewList[int]()

	list.Add(10)
	list.Add(20)

	if err := list.Set(0, 8); err != nil {
		t.Error("should have set item 0")
	}

	if err := list.Set(1, 9); err != nil {
		t.Error("should have set item 1")
	}

	if item, err := list.Get(0); err != nil || item != 8 {
		t.Error("Expected 8")
	}

	if item, err := list.Get(1); err != nil || item != 9 {
		t.Error("Expected 9")
	}
}

func Test_Insert_Empty(t *testing.T) {
	list := NewList[int]()

	if err := list.Insert(0, 99); err != nil {
		t.Error("Insert should have inserted an item")
	}

	if list.Count() != 1 {
		t.Error("should have one item in list")
	}

	if item, _ := list.Get(0); item != 99 {
		t.Error("should have got 99")
	}
}

func Test_Insert(t *testing.T) {
	list := NewList[int]()

	list.Add(10)
	list.Add(20)

	if err := list.Insert(1, 99); err != nil {
		t.Error("Insert should have inserted an item")
	}

	if list.Count() != 3 {
		t.Error("should have 3 items in list")
	}

	if item, _ := list.Get(0); item != 10 {
		t.Error("should have got 10")
	}

	if item, _ := list.Get(1); item != 99 {
		t.Error("should have got 99")
	}

	if item, _ := list.Get(2); item != 20 {
		t.Error("should have got 20")
	}
}

func Test_Insert_Bad_Index(t *testing.T) {
	list := NewList[int]()

	list.Add(10)
	list.Add(20)

	if err := list.Insert(-1, 99); err == nil {
		t.Error("Insert should have failed")
	}

	if err := list.Insert(3, 99); err == nil {
		t.Error("Insert should have failed")
	}

	if list.Count() != 2 {
		t.Error("should have 2 items in list")
	}
}

func Test_List_Stream_Empty(t *testing.T) {
	list := NewList[int]()
	stream := list.Stream()

	if stream.Iterator().MoveNext() {
		t.Error("there should be nothing to stream")
	}
}

func Test_List_Stream(t *testing.T) {
	list := NewList[int]()
	list.Add(8)

	stream := list.Stream()
	i := stream.Iterator()

	if !i.MoveNext() {
		t.Error("there should be something to stream")
	}

	if i.Current() != 8 {
		t.Error("should have 8")
	}

	if i.MoveNext() {
		t.Error("there shouldn't be any more data")
	}
}
