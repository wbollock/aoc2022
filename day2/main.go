package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// part 1 logic
	// strategy guide is outcome. determine what points outcome gives
	// A - rock  (1 pt for choosing this)
	// B - paper (2 pts for choosing)
	// C - scissors (3 pts for choosing this)
	// X-  rock
	// y - Paper
	// z - scissors
	// first column is what opp plays, second is what I play
	// win = +6
	// draw = +3
	// lose = 0

	// part 2 logic
	// X = I must lose
	// Y = I must draw
	// Z = I must win

	var totalPoints int
	dataSlice := strings.Split(string(data), "\n")
	for _, round := range dataSlice {
		// fmt.Println("jawn", round)
		roundSlice := strings.Split(round, " ")

		switch roundSlice[0] {
		case "A": // if opponent plays rock
			switch roundSlice[1] {
			case "X": // must lose, play scissors
				totalPoints += 3 // scissors lose
			case "Y": // must draw, play rock
				totalPoints += 4 // rock draw
			case "Z": // must win, paper win
				totalPoints += 8 // paper win
			}
		case "B": // paper
			switch roundSlice[1] {
			case "X":
				totalPoints += 1 // rock lose
			case "Y":
				totalPoints += 5 // paper draw
			case "Z":
				totalPoints += 9 // scissors win
			}
		case "C": // scissors
			switch roundSlice[1] {
			case "X":
				totalPoints += 2 // paper lose
			case "Y":
				totalPoints += 6 // scissors draw
			case "Z":
				totalPoints += 7 // rock win
			}

		}

	}
	fmt.Println("Total points:", totalPoints)

}
