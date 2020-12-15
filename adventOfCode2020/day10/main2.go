package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, max := parseNumberFromInput("./input.txt")
	// value := 1 + calcPath(0, max, input, 0)
	fmt.Println(input)
	diff := differences(input, max)
	fmt.Println(diff)
	countOnes(diff)

}

func countOnes(num []int) {
	comb := make(map[int]int)
	comb[1] = 1
	comb[2] = 1
	comb[3] = 2
	comb[4] = 4
	comb[5] = 7
	counter := 1
	v := 1
	i := 0
	for {
		if num[i] == 1 {
			counter++
		} else {
			v = v * comb[counter]
			counter = 1
		}
		i++

		if i == len(num) {
			v = v * comb[counter]
			counter = 1
			break
		}
	}
	fmt.Println(v)
}

func differences(num []int, max int) []int {
	sort.Ints(num)
	fmt.Println(num)
	diff := make([]int, len(num))
	last := num[0]
	for i := 1; i < len(num); i++ {
		current := num[i]
		diff[i] = current - last
		last = current
	}
	return diff[1:]
}

func parseNumberFromInput(input string) ([]int, int) {
	content, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make([]int, 1)
	numbers[0] = 0
	max := 0
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, value)
		max = int(math.Max(float64(value), float64(max)))
	}
	return numbers, max
}

func repeat(s string, times int) string {
	concat := ""
	for i := 0; i < times; i++ {
		concat += s
	}
	return concat
}
