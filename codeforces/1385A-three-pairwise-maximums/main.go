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

	max := 0
	sndMax := 0
	for _, v := range a {
		if v > max {
			sndMax = max
			max = v
		} else if v > sndMax && v != max {
			sndMax = v
		}
	}

	count := 0
	for _, v := range a {
		if v == max {
			count++
		}
	}

	if count == 2 {
		printf("YES\n")
		printf("%d %d %d\n", max, sndMax, 1)
	} else if count == 3 {
		printf("YES\n")
		printf("%d %d %d\n", max, max, max)
	} else {
		printf("NO\n")
	}

}

func main() {
	defer writer.Flush()

	var testCases int
	scanf("%d\n", &testCases)
	for i := 0; i < testCases; i++ {
		var x, y, z int
		scanf("%d %d %d", &x, &y, &z)
		// printf("%d %d %d\n", x, y, z)
		scanf("\n")
		solve([]int{x, y, z})
	}
}
