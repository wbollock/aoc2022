package main

import (
	"fmt"
	"os"
	"strings"
)

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
	for _, compartment := range inputSlice {
		var value int
		firstHalf := len(compartment) / 2

		// iterate through chars in first half of string
		for _, char := range compartment[0:firstHalf] {
			// compare each char in first half to second half of string
			result := strings.Contains(compartment[firstHalf:], string(char))

			if result {
				fmt.Println(compartment)
				// this may be stupid but we can use ASCII position to get alphabetical index
				// like if it's an uppercase char (A = 65, we subtract 64 to get index 1)
				// for lowercase, subtract 96

				if strings.ToUpper(string(char)) == string(char) {
					// this would be minus 64 from ascii index but uppercase is worth 27 through 52
					// so only minus (64-26) = 38
					value = int(char) - 38
					fmt.Println("upper case found", string(char), points, value)
				} else { // lowercase
					value = int(char) - 96
					fmt.Println("lower case found", string(char), points, value)
				}

			}

		}
		points += value
	}
	fmt.Println("Total points: ", points)
}
