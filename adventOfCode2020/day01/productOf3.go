package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func solve(a []int) {

}

func main() {
	defer writer.Flush()

	var expenseValues []int
	for {
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			break
		}
		expenseValues = append(expenseValues, i)
		// printf("%d\n", i)
	}

	mySet := map[int]bool{}

	for _, expense1 := range expenseValues {
		for _, expense2 := range expenseValues {
			mySet[expense2] = true
			diff := 2020 - (expense1 + expense2)
			_, exists := mySet[diff]
			if exists {
				printf("First number %d\n", expense1)
				printf("Second number %d\n", expense2)
				printf("Third number %d\n", diff)
				printf("Multiplication %d\n", expense1*expense2*diff)
				return
			}
		}

	}

}
