package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	numbers := parseNumbersFromInput("./input.txt")
	numberWithNoSum := findNumberWhichHasNoSum(numbers, 25)

	// numbers := parseNumbersFromInput("./test.txt")
	// numberWithNoSum := findNumberWhichHasNoSum(numbers, 5)

	if numberWithNoSum == -1 {
		panic("Hey! this never should happen. Bad input")
	}
	fmt.Printf("Number with no sum %d\n", numberWithNoSum)

	mySet := findConsequentSet(numbers, numberWithNoSum)
	w34kN355 := findEncryptionWeakn355(mySet)

	// fmt.Printf("Set %d\n", mySet)
	// acc := 0
	// for _, v := range mySet {
	// 	res := acc + v
	// 	fmt.Printf("%d + %d = %d\n", acc, v, res)
	// 	acc = res
	// }
	fmt.Printf("Encryption weakness %d\n", w34kN355)

}

func findEncryptionWeakn355(consequentSet []int) int {
	min := consequentSet[0]
	max := consequentSet[0]
	for _, c := range consequentSet {
		min = int(math.Min(float64(c), float64(min)))
		max = int(math.Max(float64(c), float64(max)))

	}
	return min + max
}

func findConsequentSet(numbers []int, numberWithNoSum int) []int {
	summatory := 0
	mySet := make([]int, 0)
	for _, number := range numbers {
		summatory += number
		mySet = append(mySet, number)

		for {
			if summatory > numberWithNoSum {
				value := mySet[0]
				summatory -= value
				// Remove first element of the set
				mySet = mySet[1:]
			} else if summatory == numberWithNoSum && len(mySet) >= 2 {
				// Found it!
				return mySet
			} else {
				break
			}
		}
	}
	return make([]int, 0)
}

func findNumberWhichHasNoSum(input []int, preableLength int) int {

	numbersInPreamble := make([]int, 0)

	for i := 0; i < preableLength; i++ {
		number := input[i]
		numbersInPreamble = append(numbersInPreamble, number)
	}

	for _, expectedResult := range input[preableLength:] {
		sumExist := findSum(numbersInPreamble, expectedResult)
		if sumExist {
			// remove first element &&
			// add new element
			numbersInPreamble = append(numbersInPreamble[1:], expectedResult)
		} else {
			return expectedResult
		}

	}
	return -1

}

func findSum(numbers []int, expectedResult int) bool {
	m := make(map[int]bool)

	for _, number := range numbers {
		m[number] = true
	}

	for i := 0; i < len(numbers); i++ {
		if _, exist := m[expectedResult-numbers[i]]; exist {
			return true
		}
	}
	return false
}

func parseNumbersFromInput(input string) []int {
	content, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers[i] = value
	}
	return numbers
}
