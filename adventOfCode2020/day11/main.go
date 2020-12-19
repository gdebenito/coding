package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	Free     = 0
	Empty    = 1
	Occupied = 2
)

func main() {
	data := readFile("./input.txt")
	matrix := generateMatrixFromString(data)
	// printMatrix(matrix)

	res := solve1(matrix)
	fmt.Println(res)

	res = solve2(matrix)
	fmt.Println(res)

}

func solve1(matrix [][]int) int {
	oldMatrix := matrix
	j := 0
	for {
		// fmt.Println("--------------")
		nextMatrix := make([][]int, len(oldMatrix))
		for i, row := range oldMatrix {
			nextMatrix[i] = make([]int, len(row))
			for j, oldState := range row {
				var newState int

				if oldState != Free {
					adj := adjacents(i, j, oldMatrix)

					if adj >= 4 && oldState == Occupied {
						newState = Empty
						// fmt.Println("Going EMPTY")
					} else if adj == 0 && oldState == Empty {
						newState = Occupied
						// fmt.Println("Going OCCUPIED")
					} else {
						newState = oldState
					}
				} else {
					newState = Free
				}

				nextMatrix[i][j] = newState
			}
		}

		if equalMatrix(oldMatrix, nextMatrix) {
			break
		}

		// fmt.Println("------old----------")
		// printMatrix(oldMatrix)
		// fmt.Println("------------------")
		// fmt.Println("")
		// fmt.Println("------new----------")
		// printMatrix(nextMatrix)
		// fmt.Println("------------------")

		j++

		oldMatrix = nextMatrix

	}

	return countInMatrix(Occupied, oldMatrix)
}

func solve2(matrix [][]int) int {
	oldMatrix := matrix
	for {
		// fmt.Println("--------------")
		nextMatrix := make([][]int, len(oldMatrix))
		for i, row := range oldMatrix {
			nextMatrix[i] = make([]int, len(row))
			for j := range row {
				newState := calculateNewState(i, j, oldMatrix)
				nextMatrix[i][j] = newState
			}
		}

		// fmt.Println("------new-----")
		// printMatrix(nextMatrix)

		if equalMatrix(oldMatrix, nextMatrix) {
			break
		}

		oldMatrix = nextMatrix

	}

	return countInMatrix(Occupied, oldMatrix)

}

func calculateNewState(i int, j int, matrix [][]int) int {
	oldState := matrix[i][j]
	var newState int

	if oldState != Free {
		adj := adjacents2(i, j, matrix)
		if adj >= 5 && oldState == Occupied {
			newState = Empty
			// fmt.Println("Going EMPTY")
		} else if adj == 0 && oldState == Empty {
			newState = Occupied
			// fmt.Println("Going OCCUPIED")
		} else {
			newState = oldState
		}
	} else {
		newState = Free
	}

	return newState
}

func findFirstSeatAt(row int, col int, matrix [][]int, vx int, vy int) int {

	nx := min(max(0, col+vx), len(matrix[0])-1)
	ny := min(max(0, row+vy), len(matrix)-1)

	if (vx != 0 && nx == col) || (vy != 0 && ny == row) {
		return 0
	}

	limitsAreReached := (nx == col && ny == row)

	if limitsAreReached {
		return 0
	}

	for {
		value := matrix[ny][nx]

		if value == Occupied {
			return 1
		} else if value == Empty {
			return 0
		}

		nx += vx
		ny += vy

		if nx > len(matrix[0])-1 ||
			ny > len(matrix)-1 ||
			nx < 0 ||
			ny < 0 {
			return 0
		}
	}

}

func adjacents2(row int, col int, matrix [][]int) int {
	adj := 0

	adj += findFirstSeatAt(row, col, matrix, 0, 1)
	adj += findFirstSeatAt(row, col, matrix, 1, 0)
	adj += findFirstSeatAt(row, col, matrix, 0, -1)
	adj += findFirstSeatAt(row, col, matrix, -1, 0)
	adj += findFirstSeatAt(row, col, matrix, 1, 1)
	adj += findFirstSeatAt(row, col, matrix, 1, -1)
	adj += findFirstSeatAt(row, col, matrix, -1, 1)
	adj += findFirstSeatAt(row, col, matrix, -1, -1)
	// fmt.Printf("Adjuntos de: %d %d = %d\n", row, col, adj)

	return adj

}

func adjacents(row int, col int, matrix [][]int) int {
	fromCol := max(0, col-1)
	toCol := min(col+1, len(matrix)-1)

	fromRow := max(0, row-1)
	toRow := min(row+1, len(matrix[0])-1)

	adj := 0

	// fmt.Printf("From Row: %d to %d, col %d to %d \n", fromRow, toRow, fromCol, toCol)

	for i := fromRow; i <= toRow; i++ {
		for j := fromCol; j <= toCol; j++ {
			// fmt.Printf("%d ", matrix[i][j])
			if j != col || i != row {
				// fmt.Printf("  i:%d j:%d,v:%d\n", i, j, matrix[i][j])
				if matrix[i][j] == Occupied {
					adj++
				}
			}
		}
		// fmt.Printf("\n")
	}

	// fmt.Printf("Occupied: %d\n", adj)

	return adj

}

func countInMatrix(v int, matrix [][]int) int {
	res := 0

	for _, row := range matrix {
		for _, cell := range row {

			if cell == v {
				res++
			}

		}
	}

	return res

}
func printMatrix(matrix [][]int) {

	for _, row := range matrix {
		for _, cell := range row {
			switch cell {
			case Free:
				fmt.Printf(".")
			case Empty:
				fmt.Printf("L")
			case Occupied:
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func equalMatrix(a [][]int, b [][]int) bool {
	for i, row := range a {
		for j, cell := range row {
			if b[i][j] != cell {
				return false
			}
		}
	}
	return true
}

func generateMatrixFromString(data string) [][]int {
	lines := strings.Split(data, "\n")

	matrix := make([][]int, len(lines))

	for i, line := range lines {
		matrix[i] = make([]int, len(line))
		for j, char := range line {
			if char == 'L' {
				matrix[i][j] = Empty
			} else {
				matrix[i][j] = Free
			}
		}
	}

	return matrix
}

func readFile(name string) string {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(contents)
}
