package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	heights := readContents("./input.txt")

	fmt.Println(solve1(heights))
	fmt.Println(solve2(heights))

}

func solve1(heights []uint) uint {
	var last uint = 0
	var countIncreased uint = 0

	for i, actual := range heights {
		if i == 0 {
			last = actual
			continue
		}

		if actual > last {
			countIncreased = countIncreased + 1
		}

		last = actual
	}

	return countIncreased
}

func solve2(heights []uint) uint {

	var window []uint
	var countIncreased uint = 0

	for i, height := range heights {

		if i < 3 {
			window = append(window,height)
			continue
		}

		lastSum := sumSlice(window)
		window = append(window[1:], height)
		nextSum := sumSlice(window)

		if nextSum > lastSum {
			countIncreased = countIncreased + 1
		}

	}

	return countIncreased

}

func sumSlice(slice []uint) uint {
	var sum uint = 0
	for _, i := range slice {
		sum = sum + i
	}
	return sum
}

func readContents(filename string) ([]uint) {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	heightsRaw := strings.Split(string(content), "\n")

	heights := make([]uint, 0)
	for _, heightRaw := range heightsRaw {

		height, err := strconv.Atoi(heightRaw)

		if err != nil {
			continue
		}

		heights = append(heights, uint(height))
	}

	return heights

}
