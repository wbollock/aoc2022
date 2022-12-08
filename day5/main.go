package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func splitByEmptyNewline(str string) []string {
	// shamelessly stolen from https://stackoverflow.com/a/36958185
	// but I had the idea!
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(str, "\n")

	return regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)

}

func prettyMatrixDisplay(crateMatrix [][]string) {

	for row := 0; row < len(crateMatrix); row++ {

		for column := 0; column < len(crateMatrix[0]); column++ {
			fmt.Print(crateMatrix[row][column])
		}
		fmt.Print("\n")
	}

}

func partOne(input []byte) (topCrates string) {
	// first step:
	// take in crates and their order as variable
	// take in orders as separate variable

	// empty new line seperates the crates and the instructions
	inputSlice := splitByEmptyNewline(string(input))

	crateString := inputSlice[0]
	orderString := inputSlice[1]

	// get crateString into useable format by removing [ ]
	// crateString = strings.Replace(crateString, "[", "", -1)
	// crateString = strings.Replace(crateString, "]", "", -1)

	// fmt.Println("crates:\n", crateString)
	// crateMatrix := make([][]string, 1)
	var crateMatrix [][]string

	// for every row of crates
	for _, crate := range strings.Split(crateString, "\n") {
		// crateMatrix[counter] = append(crateMatrix[0], jawn)

		var spaceCounter int
		var row []string
		for _, char := range jawn {
			// fmt.Println("CHAR!", string(char))

			if string(char) == " " {
				spaceCounter++
			} else {
				spaceCounter = 0
			}

			// if we have four spaces in a row we have no crate, so we can safely append a space
			if spaceCounter == 4 {
				row = append(row, " ")
			} else if (string(char) != " ") && (string(char) != "[") && (string(char) != "]") {
				// fmt.Println("appending!", string(char))
				row = append(row, string(char))
			}
		}
		crateMatrix = append(crateMatrix, row)
		// fmt.Println("row:", row)
		// fmt.Println("counter:", counter)

	}

	// second step:
	// perform actions in orders
	// fmt.Println("orders:", orderString)

	prettyMatrixDisplay(crateMatrix)
	for _, order := range strings.Split(orderString, "\n") {
		re := regexp.MustCompile("[0-9]+")
		orderNumbers := re.FindAllString(order, -1)
		// always have 3 numbers in order
		// first number is how many crates down to move (orderNumbers[0])

		// second number is source column (orderNumbers[1])

		// third numbers is destination column (orderNumbers[2])
		fmt.Println("move ", orderNumbers[0], "crates from column", orderNumbers[1], "to column", orderNumbers[2])

		cratesToMove, err := strconv.Atoi(orderNumbers[0])
		if err != nil {
			panic(err)
		}

		sourceColumn, err := strconv.Atoi(orderNumbers[1])
		if err != nil {
			panic(err)
		}

		destColumn, err := strconv.Atoi(orderNumbers[2])
		if err != nil {
			panic(err)
		}

		// for counter := 0; counter <= cratesToMove; counter++ {
		// for row := 0; row < len(crateMatrix[sourceColumn-1]); row++ {

		// 	// crateToGrab is going to be top char in sourceColumn
		// 	var crateToGrab string
		// 	if crateMatrix[row][sourceColumn-1] != " " {
		// 		crateToGrab = crateMatrix[row][sourceColumn-1]
		// 		fmt.Println("crate to grab!", crateToGrab, sourceColumn-1, row)
		// 	}
		// 	crateMatrix[row] = append([]string{crateToGrab}, crateMatrix[row]...)
		// 	fmt.Println(crateMatrix[row])

		// // if we don't have a blank space at the destination, lets make one
		// if crateMatrix[destColumn-1][0] != " " {
		// 	blankSlice := []string{" "}
		// 	crateMatrix[destColumn-1] = append(blankSlice, crateMatrix[destColumn-1]...)
		// }
		// // if we got a blank space it's easier! now we should have one
		// if crateMatrix[destColumn-1][0] == " " {
		// 	crateMatrix[destColumn-1][0] = crateToGrab
		// 	crateMatrix[row][sourceColumn-1] = " "
		// 	break
		// }
		// 	}
		// }

		for counter := 0; counter <= cratesToMove; counter++ {

			// crateToGrab := crateMatrix[counter][sourceColumn-1]
			var crateToGrab string
			for row := 0; row < len(crateMatrix[sourceColumn-1]); row++ {
				if crateMatrix[row][sourceColumn-1] != " " {
					crateToGrab = crateMatrix[row][sourceColumn-1]
					fmt.Println("crate to grab!", crateToGrab, sourceColumn-1, row)
				}
			}

			// if we don't have a blank space at the destination, lets make one
			if crateMatrix[destColumn-1][0] != " " {
				blankSlice := []string{" "}
				crateMatrix[destColumn-1] = append(blankSlice, crateMatrix[destColumn-1]...)
			}
			fmt.Println("")
			fmt.Println("appended blank slice")
			prettyMatrixDisplay(crateMatrix)
			fmt.Println("")
			// if we got a blank space it's easier! now we should have one
			if crateMatrix[destColumn-1][0] == " " {
				crateMatrix[destColumn-1][0] = crateToGrab
				crateMatrix[counter][sourceColumn-1] = " "
				break
			}
		}

		// fmt.Println(crateMatrix[0][destColumn-1])
		// crateMatrix[destColumn-1] = append([]string{crateToGrab}, crateMatrix[destColumn-1]...)

		// crateMatrix[sourceColumn][counter] = " "

		// fmt.Println("Order done! displaying")
		prettyMatrixDisplay(crateMatrix)
		// fmt.Println("done display")
	}

	// fmt.Println("JAWN:", crateMatrix[0][1])

	// third step:
	// gather crates at top of pile and return them

	// iterate my *column* first
	// get the first char in each column
	for row := 0; row < len(crateMatrix[0]); row++ {

		for column := 0; column < len(crateMatrix[row]); column++ {
			if crateMatrix[column][row] != " " {
				topCrates = topCrates + crateMatrix[column][row]
				break
			}

		}
	}

	return topCrates

}

func main() {

	input := readInput("input")
	answer := partOne(input)
	fmt.Println("TOP CRATES:", answer)
}
