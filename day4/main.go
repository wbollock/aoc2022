package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// this is nasty but a way to iterate through our elf strings and see if they have any matches at all
func compareElfs(biggerElfRange string, smallerElfRange string) (matchFound bool) {

	elfChars := strings.Split(smallerElfRange, ",")

	for _, char := range elfChars {
		if char == "" {
			break
		}
		if strings.Contains(biggerElfRange, char) {
			matchFound = true
			break
		}
	}

	return matchFound

}

func main() {
	// read file
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// part 1 logic
	// pairs 1-3,4-6 are {1,2,3},{4,5,6}
	// need to count how many pairs fully envelope each other

	// first explode both pairs into slices of strings (so we can use strings.Contains() later)
	// then see if either slice fully contains the other

	inputSlice := strings.Split(string(input), "\n")

	var matches int

	for _, elfPair := range inputSlice {
		elfSlice := strings.Split(elfPair, ",")

		var elfRange []string

		for _, elf := range elfSlice { // divide elf pair into individual strings of numbers
			// e.g 2-4 becomes 2,3,4
			individualElfSlice := strings.Split(elf, "-")

			lowElfRange, err := strconv.Atoi(individualElfSlice[0])
			if err != nil {
				panic(err)
			}

			highElfRange, err := strconv.Atoi(individualElfSlice[1])
			if err != nil {
				panic(err)
			}

			individualElfRange := make([]int, highElfRange-lowElfRange+1)
			// iterate our 1-5 pair into 12345
			for i := range individualElfRange {
				individualElfRange[i] = i + lowElfRange
			}

			individualElfRangeStrings, _ := json.Marshal(individualElfRange)

			cleanStrings := strings.Trim(string(individualElfRangeStrings), "[]")

			// this is so ass I hope no one looks at this
			// my approach sucked here using strings. edge case for single digits, need 0 appended
			preZeroAppendStrings := strings.Split(cleanStrings, ",")
			var postZeroAppendStrings string

			for _, char := range preZeroAppendStrings {

				// because if we have a single digit "9" it could match to 99, 79, etc
				// so turn "9" into "09"
				if len(char) == 1 {
					char = "0" + char
				}
				char = char + ","
				postZeroAppendStrings += char
			}

			elfRange = append(elfRange, postZeroAppendStrings)

		}

		var matchFound bool
		if len(elfRange[0]) > len(elfRange[1]) {

			matchFound = compareElfs(elfRange[0], elfRange[1])
			if matchFound {
				matches++
			}

		} else if len(elfRange[1]) > len(elfRange[0]) {

			matchFound = compareElfs(elfRange[0], elfRange[1])
			if matchFound {
				matches++
			}

		} else { // equal length pairs
			if elfRange[0] == elfRange[1] {
				matches++
			} else {
				matchFound = compareElfs(elfRange[0], elfRange[1])
				if matchFound {
					matches++
				}
			}
		}

	}
	fmt.Println("Total matches:", matches)
}
