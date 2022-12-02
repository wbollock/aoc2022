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

	var totalPoints int
	dataSlice := strings.Split(string(data), "\n")
	for _, round := range dataSlice {
		// fmt.Println("jawn", round)
		roundSlice := strings.Split(round, " ")

		switch roundSlice[0] {
		case "A": // if opponent plays rock
			switch roundSlice[1] {
			case "X":
				totalPoints += 4 // rock draw
			case "Y":
				totalPoints += 8 // paper win
			case "Z":
				totalPoints += 3 // scissors lose
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
				totalPoints += 7 // rock win
			case "Y":
				totalPoints += 2 // paper lose
			case "Z":
				totalPoints += 6 // scissors draw
			}

		}

	}
	fmt.Println("Total points:", totalPoints)

}
