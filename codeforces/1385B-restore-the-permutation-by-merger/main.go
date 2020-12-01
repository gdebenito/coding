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

func solve(p []int, length int) {
	m := make(map[int]bool)
	i := 0

	for _, v := range p {
		if m[v] == false {
			m[v] = true
			if i != 0 {
				printf(" ")
			}
			printf("%d", v)

			i++
		}
	}
	printf("\n")

}

func main() {
	defer writer.Flush()

	var testCases int
	scanf("%d\n", &testCases)
	for i := 0; i < testCases; i++ {
		var length int
		scanf("%d\n", &length)
		var lengthOfTestCase = length * 2
		var p = make([]int, lengthOfTestCase)
		for j := 0; j < lengthOfTestCase; j++ {
			if j < lengthOfTestCase-1 {
				scanf("%d ", &p[j])
			} else {
				scanf("%d\n", &p[j])
			}
		}
		solve(p, length)
	}
}
