package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func readInput(file string) (input []byte) {
	// read file
	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return input
}

func partOne(input []byte) (topCrates string) {
	// first step:
	// take in crates and their order as variable
	// take in orders as separate variable

	split := strings.Split(string(input), "\n\n")
	crates := strings.Split(split[0], "\n")
	keys := crates[len(crates)-1]

	stack := map[rune]string{}

	for _, crate := range crates {
		for index, char := range crate {
			if unicode.IsLetter(char) {
				stack[[]rune(keys)[index]] += string(char)
			}
		}
	}

	// second step:
	// perform actions in orders

	for _, order := range strings.Split(strings.TrimSpace(split[1]), "\n") {
		var quantity int
		var from rune
		var to rune

		// this is neat
		fmt.Sscanf(order, "move %d from %c to %c", &quantity, &from, &to)

		for i := 0; i < quantity; i++ {
			stack[to] = stack[from][:1] + stack[to]
			stack[from] = stack[from][1:]
		}
	}

	var part1 string

	for _, char := range strings.ReplaceAll(keys, " ", "") {
		part1 += stack[char][:1]
	}
	return part1

}

func main() {

	input := readInput("input")
	answer := partOne(input)
	fmt.Println("TOP CRATES:", answer)
}
