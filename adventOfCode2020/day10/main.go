package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	input, max := parseNumbersFromInput("./input.txt")
	// input, max := parseNumbersFromInput("./test.txt")

	timesOne := 0
	timesTwo := 0
	timesThree := 0

	currentValue := 0

	for currentValue < max {

		if _, isInMap := input[currentValue+1]; isInMap {
			fmt.Printf("%d + 1\n", currentValue)
			currentValue++
			timesOne++
		} else if _, isInMap := input[currentValue+2]; isInMap {
			fmt.Printf("%d + 2\n", currentValue)
			currentValue += 2
			timesTwo++
		} else {
			fmt.Printf("%d + 3\n", currentValue)
			currentValue += 3
			timesThree++
		}

	}
	timesThree++
	currentValue += 3

	fmt.Printf("Times one %d\n", timesOne)
	fmt.Printf("Times three %d\n", timesThree)
	fmt.Printf("Mult %d\n", timesOne*timesThree)

}

func parseNumbersFromInput(input string) (map[int]bool, int) {
	content, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make(map[int]bool)
	max := 0
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers[value] = true
		max = int(math.Max(float64(value), float64(max)))
	}
	return numbers, max
}
