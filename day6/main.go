package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(file string) (input []byte) {
	// read file
	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return input
}

func partOne(input []byte) (partone int) {

	// take in input
	inputSlice := strings.Split(string(input), "\n")

	for _, message := range inputSlice {
		// would be best to parse each string as slice of runes? is that a thing in Go?
		// yea

		runeSlice := []rune(message)
		// iterate through our message
		for counter := 0; counter < len(runeSlice); counter++ {

			var badMessage bool
			var marker int

			// create mini string of 4 characters
			upperBound := counter + 4
			if upperBound > len(runeSlice) {
				upperBound = len(runeSlice) // hard limit
			}

			miniString := runeSlice[counter:upperBound]

			// for every string in our miniString
			for _, char := range miniString {
				count := strings.Count(string(miniString), string(char))

				// if we detect a character more than once, our message is bunk and we must start over with miniString + shifted one right
				if count != 1 {
					badMessage = true
					break
				} else {
					marker = upperBound
				}
			}

			if !badMessage {
				partone = marker
				break
			}

		}

	}

	return partone
}

func partTwo(input []byte) (partone int) {

	// take in input
	inputSlice := strings.Split(string(input), "\n")

	for _, message := range inputSlice {
		// would be best to parse each string as slice of runes? is that a thing in Go?
		// yea

		runeSlice := []rune(message)
		// iterate through our message
		for counter := 0; counter < len(runeSlice); counter++ {

			var badMessage bool
			var marker int

			// create mini string of 4 characters
			upperBound := counter + 14
			if upperBound > len(runeSlice) {
				upperBound = len(runeSlice) // hard limit
			}

			miniString := runeSlice[counter:upperBound]

			// for every string in our miniString
			for _, char := range miniString {
				count := strings.Count(string(miniString), string(char))

				// if we detect a character more than once, our message is bunk and we must start over with miniString + shifted one right
				if count != 1 {
					badMessage = true
					break
				} else {
					marker = upperBound
				}
			}

			if !badMessage {
				partone = marker
				break
			}

		}

	}

	return partone
}

func main() {

	input := readInput("input")
	partoneAnswer := partOne(input)
	parttwoAnswer := partTwo(input)
	fmt.Println("Part 1", partoneAnswer)
	fmt.Println("Part 2", parttwoAnswer)
}
