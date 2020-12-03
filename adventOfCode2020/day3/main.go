package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	trees := []int{goThroughTheSlope(1, 1),
		goThroughTheSlope(3, 1),
		goThroughTheSlope(5, 1),
		goThroughTheSlope(7, 1),
		goThroughTheSlope(1, 2)}

	totalTrees := 1

	for _, numTrees := range trees {
		fmt.Println(strconv.Itoa(numTrees))
		totalTrees = totalTrees * numTrees
	}

	fmt.Printf("Multiplication of all trees together: %d\n", totalTrees)

}

// Velocity is the right steps
// Gravity is the down steps
func goThroughTheSlope(velocity int, gravity int) int {

	tree := byte('#')
	numberOfTrees := 0
	numberOfSquares := 0
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	width := 31
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("%s\n", line)

		if line[(i*velocity)%width] == tree {
			numberOfTrees++
		} else {
			numberOfSquares++
		}

		i++

		if gravity == 2 {
			scanner.Scan()
		}

	}

	// fmt.Printf("Number of trees %d\n", numberOfTrees)
	// fmt.Printf("Number of squares %d\n", numberOfSquares)

	return numberOfTrees

}
