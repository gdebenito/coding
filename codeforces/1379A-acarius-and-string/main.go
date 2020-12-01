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

func hasSubstring(s string, pattern string) {

}

func solve(s string, n int) {

}

func main() {
	defer writer.Flush()

	var testCases int
	scanf("%d\n", &testCases)
	for i := 0; i < testCases; i++ {
		var length int
		var test string

		scanf("%d\n", &length)
		scanf("%s\n", &test)
		solve(test, length)
		// printf("%s\n", test)
	}
}
