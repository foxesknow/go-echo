package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Chunk_Empty(t *testing.T) {
	numbers := data.EmptyStream[int]()
	chunks := Chunk(numbers, 2)
	slices := ToSlice(chunks)

	if len(slices) != 0 {
		t.Error("There shouldn't be any chunks")
	}
}

func Test_Chunk(t *testing.T) {
	numbers := data.FromValues(5, 6, 10, 11, 14, 15)
	chunks := Chunk(numbers, 3)
	slices := ToSlice(chunks)

	if len(slices) != 2 {
		t.Error("There should be 3 chunks")
	}

	chunk1 := slices[0]
	if !(chunk1[0] == 5 && chunk1[1] == 6 && chunk1[2] == 10) {
		t.Error("invalid values in first chunk")
	}

	chunk2 := slices[1]
	if !(chunk2[0] == 11 && chunk2[1] == 14 && chunk2[2] == 15) {
		t.Error("invalid values in first chunk")
	}
}

func Test_Chunk_Last_Not_Full(t *testing.T) {
	numbers := data.FromValues(5, 6, 10, 11, 14)
	chunks := Chunk(numbers, 3)
	slices := ToSlice(chunks)

	if len(slices) != 2 {
		t.Error("There should be 3 chunks")
	}

	chunk1 := slices[0]
	if !(chunk1[0] == 5 && chunk1[1] == 6 && chunk1[2] == 10) {
		t.Error("invalid values in first chunk")
	}

	chunk2 := slices[1]
	if len(chunk2) != 2 {
		t.Error("last chuck should only have 2 items")
	}

	if !(chunk2[0] == 11 && chunk2[1] == 14) {
		t.Error("invalid values in first chunk")
	}
}
