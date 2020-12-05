package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	max := 0
	ids := make(map[int]bool)
	for scanner.Scan() {
		row := 0
		col := 0
		i := 64
		j := 4
		line := scanner.Text()

		frontOrBack := line[0:7]

		for _, fOrB := range frontOrBack {

			if fOrB == 'B' {
				row += i
			}

			i = i / 2

		}

		leftOrRight := line[7:10]

		for _, lOrR := range leftOrRight {

			if lOrR == 'R' {
				col += j
			}

			j = j / 2

		}

		value := row*8 + col
		if value > max {
			max = value
		}
		ids[value] = true
		// fmt.Printf("value: %d\n", value)

	}
	fmt.Printf("max value: %d\n", max)

	fmt.Printf("Missing ids: ")
	for i := 0; i < 888; i++ {
		if _, inMap := ids[i]; !inMap {
			fmt.Printf("%d,", i)
		}

	}
	fmt.Printf("\n")

}
