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

func main() {
	defer writer.Flush()

	var testCases int
	scanf("%d\n", &testCases)
	for i := 0; i < testCases; i++ {
		var digits int
		scanf("%d\n", &digits)
		var data string
		scanf("%s\n", &data)
		res := solve(data)
		if res {
			printf("YES\n")
		} else {
			printf("NO\n")
		}
	}
}

func solve(a string) bool {
	if a[0:4] == "2020" {
		return true
	}
	l := len(a)
	for i := 0; i <= 4; i++ {
		arr := ""
		arr += a[0 : 4-i]
		arr += a[l-i:]
		if arr == "2020" {
			return true
		}
	}

	return false
}
