package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	dataSlice := strings.Split(string(data), "\n")
	var sum, highestSum int
	topElvesSlice := []int{0, 0, 0} // initialize with 3 junk values to find top three highest elves

	for _, elf := range dataSlice {
		if elf != "" { // new line separated
			elfInt, err := strconv.Atoi(elf)
			if err != nil {
				panic(err)
			}
			sum = sum + elfInt

		} else {
			var counter, smallestElfPosition int
			potentialSmallestElf := topElvesSlice[0] // just test if the first slice element is the smallest

			for _, topElves := range topElvesSlice[1:] {
				counter++
				if topElves < potentialSmallestElf {
					potentialSmallestElf = topElves
					smallestElfPosition = counter
					fmt.Println("Smallest current elf:", topElves, "position", smallestElfPosition)
				}
			}
			if sum > potentialSmallestElf {
				fmt.Println("new higher sum found!", highestSum)
				highestSum = sum
				topElvesSlice[smallestElfPosition] = highestSum

				fmt.Println("Top elves:", topElvesSlice)
			}
			sum = 0
			counter = 0
			smallestElfPosition = 0

		}
	}

	var totalElfCalories int
	for _, elves := range topElvesSlice {
		totalElfCalories += elves
	}
	fmt.Println("done", totalElfCalories)
}
