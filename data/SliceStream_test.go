package data

import (
	"testing"
)

func Test_FromValues(t *testing.T) {
	stream := FromValues("hello", "goodbye")
	i := stream.GetStream()

	if !i.MoveNext() {
		t.Error("should have moved")
	}

	if i.Current() != "hello" {
		t.Error("should have hello")
	}

	if !i.MoveNext() {
		t.Error("should have moved")
	}

	if i.Current() != "goodbye" {
		t.Error("should have goodbye")
	}

	if i.MoveNext() {
		t.Error("should not have moved")
	}

	collection, ok := stream.(Collection)
	if !ok {
		t.Error("stream must implement Collection")
	}

	if collection.Count() != 2 {
		t.Error("stream should have count of 2")
	}
}

func Test_FromSlice(t *testing.T) {
	stream := FromSlice([]string{"hello", "goodbye"})
	i := stream.GetStream()

	if !i.MoveNext() {
		t.Error("should have moved")
	}

	if i.Current() != "hello" {
		t.Error("should have hello")
	}

	if !i.MoveNext() {
		t.Error("should have moved")
	}

	if i.Current() != "goodbye" {
		t.Error("should have goodbye")
	}

	if i.MoveNext() {
		t.Error("should not have moved")
	}

	collection, ok := stream.(Collection)
	if !ok {
		t.Error("stream must implement Collection")
	}

	if collection.Count() != 2 {
		t.Error("stream should have count of 2")
	}
}

func Test_FromSlice_Indexable(t *testing.T) {
	stream := FromSlice([]string{"hello", "goodbye"})
	i := stream.GetStream()

	if !i.MoveNext() {
		t.Error("should have moved")
	}

	if i.Current() != "hello" {
		t.Error("should have hello")
	}

	if !i.MoveNext() {
		t.Error("should have moved")
	}

	if i.Current() != "goodbye" {
		t.Error("should have goodbye")
	}

	if i.MoveNext() {
		t.Error("should not have moved")
	}

	indexable, ok := stream.(IndexableCollection[string])
	if !ok {
		t.Error("stream must implement IndexableCollection")
	}

	if indexable.Count() != 2 {
		t.Error("stream should have count of 2")
	}

	if item, err := indexable.Get(0); item != "hello" || err != nil {
		t.Error("failed to get item 0")
	}

	if item, err := indexable.Get(1); item != "goodbye" || err != nil {
		t.Error("failed to get item 1")
	}

	if _, err := indexable.Get(3); err == nil {
		t.Error("index was invalid")
	}
}
