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
		var lengthArr int
		scanf("%d\n", &lengthArr)
		arr := make([]int, lengthArr)
		for i := 0; i < lengthArr; i++ {
			var v int
			if i == 0 {
				scanf("%d", &v)
			} else {
				scanf(" %d", &v)
			}
			arr[i] = v
		}
		res := solve(arr)
		for i, v := range res {
			if i == 0 {
				printf("%d", v)
			} else {
				printf(" %d", v)
			}
		}
		printf("\n")
		scanf("\n")
	}
}

func solve(arr []int) []int {

	length := len(arr)
	halfLength := len(arr) / 2
	isPair := length % 2
	if isPair != 0 {
		halfLength++
	}
	res := make([]int, length)

	// fmt.Println(arr)
	// fmt.Println(length)
	// fmt.Println(halfLength)

	for i, elem := range arr {
		if i < halfLength {
			res[i*2] = elem
		} else {
			res[length-1-2*(i-halfLength)-isPair] = elem
		}
	}

	return res

}
