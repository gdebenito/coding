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

	for _, expense := range expenseValues {
		mySet[expense] = true
		diff := 2020 - expense
		_, exists := mySet[diff]
		if exists {
			printf("First number %d\n", expense)
			printf("Second number %d\n", diff)
			printf("Multiplication %d\n", expense*diff)
			break
		}

	}
}
