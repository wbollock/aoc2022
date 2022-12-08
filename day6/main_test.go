package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {

	input := readInput("testinput")

	got := partOne(input)
	want := 7

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	} else {
		t.Logf("got %v, wanted %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {

	input := readInput("testinput")

	got := partTwo(input)
	want := 19

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	} else {
		t.Logf("got %v, wanted %v", got, want)
	}
}
