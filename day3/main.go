package main

import (
	"fmt"
	"os"
	"strings"
)

func getMatchValue(char rune) (value int) {

	// this may be stupid but we can use ASCII position to get alphabetical index
	// like if it's an uppercase char (A = 65, we subtract 64 to get index 1)
	// for lowercase, subtract 96

	if strings.ToUpper(string(char)) == string(char) {
		// this would be minus 64 from ascii index but uppercase is worth 27 through 52
		// so only minus (64-26) = 38
		value = int(char) - 38
	} else { // lowercase
		value = int(char) - 96
	}
	return value

}

func main() {
	// read file
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// part 1 logic
	// first half of string is one compartment, second half is another
	// compare matching characters between compartments. will be one lowercase or uppercase match

	inputSlice := strings.Split(string(input), "\n")
	var points int
	for counter := 0; counter <= len(inputSlice); counter += 3 {
		var value int
		if counter < len(inputSlice) {

			// iterate through chars in first string, compare to second. must have match there to continue
			for _, char := range inputSlice[counter] {
				// compare each char in first half to second half of string
				result := strings.Contains(inputSlice[counter+1], string(char))

				if result {
					// we then need to see if our char is also in the last group member
					result = strings.Contains(inputSlice[counter+2], string(char))

					if result {
						value = getMatchValue(char)
					}
				}
			}
		}
		points += value
	}
	fmt.Println("Total points: ", points)
}
