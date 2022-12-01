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
	var sum int
	topElvesSlice := []int{0, 0, 0} // initialize with 3 junk values to find top three highest elves

	for _, elf := range dataSlice { // iterate through all elves (lines all together until a newline)
		if elf != "" { // new line separated
			elfInt, err := strconv.Atoi(elf)
			if err != nil {
				panic(err)
			}
			sum = sum + elfInt // total calories for one elf

		} else {
			var counter, smallestElfPosition int
			smallestElf := topElvesSlice[0] // just test if the first slice element is the smallest

			for _, topElves := range topElvesSlice[1:] {
				counter++
				if topElves < smallestElf {
					smallestElf = topElves
					smallestElfPosition = counter
				}
			}
			if sum > smallestElf {
				topElvesSlice[smallestElfPosition] = sum
				fmt.Println("Top elves:", topElvesSlice)
			}
			sum = 0
			counter = 0
		}
	}

	var totalElfCalories int
	for _, elves := range topElvesSlice {
		totalElfCalories += elves
	}
	fmt.Println("done", totalElfCalories)
}
