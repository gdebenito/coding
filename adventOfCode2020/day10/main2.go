package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Graph struct {
	length            int
	possibleJumps     map[int][]int
	possibleJumpsBack map[int][]int
}

func main() {
	input, max := parseNumbersFromInput("./test2.txt")
	// value := 1 + calcPath(0, max, input, 0)
	input[0] = true
	g := generateGraph(input)
	fmt.Println(g.possibleJumps)
	fmt.Println(g.possibleJumpsBack)
	fmt.Println(max)
	count := g.countPaths(0, max, 0)

	fmt.Println(count)
}

func generateGraph(input map[int]bool) *Graph {
	g := Graph{
		length:            len(input),
		possibleJumps:     make(map[int][]int),
		possibleJumpsBack: make(map[int][]int),
	}
	for k := range input {
		g.possibleJumpsBack[k] = make([]int, 0)
	}

	for k := range input {
		arr := make([]int, 0)

		for i := 1; i <= 3; i++ {

			if _, isInMap := input[k+i]; isInMap {
				arr = append(arr, k+i)

				aux := g.possibleJumpsBack[k+i]
				aux = append(aux, k)
				g.possibleJumpsBack[k+i] = aux
			}

		}

		g.possibleJumps[k] = arr
	}

	return &g

}

func (g *Graph) countPaths(current int, destination int, count int) int {
	if current == destination {
		// fmt.Printf("%d %d", current, destination)
		count++
	} else {
		possibleJumps := g.possibleJumps[current]
		// fmt.Println(possibleJumps)
		for _, nextJump := range possibleJumps {
			// fmt.Printf("%d %d\n", nextJump, destination)
			count = g.countPaths(nextJump, destination, count)
		}
	}
	return count
}

func parseNumbersFromInput(input string) (map[int]bool, int) {
	content, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make(map[int]bool)
	max := 0
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers[value] = true
		max = int(math.Max(float64(value), float64(max)))
	}
	return numbers, max
}

func repeat(s string, times int) string {
	concat := ""
	for i := 0; i < times; i++ {
		concat += s
	}
	return concat
}
