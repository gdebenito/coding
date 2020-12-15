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

	var testCases int
	scanf("%d\n", &testCases)
	for i := 0; i < testCases; i++ {
		// var x, y, z int
		// scanf("%d %d %d", &x, &y, &z)
		// printf("%d %d %d\n", x, y, z)
	}
}
