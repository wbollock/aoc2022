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
	var counter, sum, highestSum, highestSumPosition int

	for _, element := range dataSlice {
		counter++
		if element != "" {
			// fmt.Println(element)
			elementInt, err := strconv.Atoi(element)
			if err != nil {
				panic(err)
			}
			sum = sum + elementInt
		} else {
			if sum > highestSum {
				fmt.Println("new highest sum found!", highestSum)
				highestSum = sum
				highestSumPosition = counter
			}
			sum = 0
		}

	}

	fmt.Println("done", highestSum, highestSumPosition)
}
