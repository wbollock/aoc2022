package main

import (
	"testing"
)

var testInput string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPartTwo(t *testing.T) {

	input := []byte(testInput)

	got := partTwo(input)
	want := int(4)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("got %d %d", got, want)
	}
}
