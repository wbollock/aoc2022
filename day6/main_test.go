package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {

	input := readInput("testinput")

	got := partOne(input)
	want := 11

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	} else {
		t.Logf("got %v, wanted %v", got, want)
	}
}
