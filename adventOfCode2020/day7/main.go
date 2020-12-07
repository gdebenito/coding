package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Bag struct {
	Repeats int
	Name    string
}

func main() {
	contents, err := ioutil.ReadFile("./input.txt")
	check(err)

	lines := strings.Split(string(contents), "\n")

	bagsContainedInBag := make(map[string][]Bag)
	inversedBagGraph := make(map[string]map[string]bool)
	shinyBags := make(map[string]bool)

	for _, line := range lines {

		if line == "" {
			break
		}

		bagAndContains := strings.Split(line, " bags contain ")
		bag := bagAndContains[0]
		contains := strings.Split(bagAndContains[1], ",")

		var bagsContained []Bag

		for _, contain := range contains {
			if contain == "no other bags." {
				continue
			}
			trimmed := trim(contain, []string{" ", ".", "bags", "bag"})
			containingBag := trim(trimmed[2:], []string{" "})
			_, hasData := inversedBagGraph[containingBag]
			if !hasData {
				inversedBagGraph[containingBag] = make(map[string]bool)
			}
			myMap := inversedBagGraph[containingBag]
			myMap[bag] = true

			number, err := strconv.Atoi(trimmed[0:1])
			check(err)
			repeats := math.Max(float64(number), 1)
			bagsContained = append(bagsContained, Bag{
				Repeats: int(repeats),
				Name:    containingBag,
			})
			if containingBag == "shiny gold" {
				shinyBags[bag] = true
			}
		}
		bagsContainedInBag[bag] = bagsContained
	}

	// Las bolsas que pueden contener shiny bags
	resultBags := make(map[string]bool)
	for k := range shinyBags {
		// fmt.Println(k)
		resultBags[k] = true
		search(k, inversedBagGraph, resultBags, 0)
	}

	// fmt.Println(resultBags)
	fmt.Printf("Bags can contain shiny gold bag: %d\n", len(resultBags))

	total := sumBags("shiny gold", bagsContainedInBag, 0) - 1
	fmt.Printf("Shiny gold bag must contain %d bags\n", total)

}

func sumBags(key string, bagGraph map[string][]Bag, level int) int {
	bags, hasData := bagGraph[key]
	sum := 1
	if !hasData {
		// Empty bag only contains one bag (self)
		return sum
	}

	for _, bag := range bags {
		// fmt.Println(repeat(" ", level) + key + " = " + strconv.Itoa(bag.Repeats) + " * ")
		value := sumBags(bag.Name, bagGraph, level+1)
		sum += bag.Repeats * value
		//fmt.Println(repeat(" ", level) + strconv.Itoa(value))
	}
	return sum
}

func search(key string, bagGraph map[string]map[string]bool, resultBags map[string]bool, level int) {
	data, hasData := bagGraph[key]
	if !hasData {
		return
	}

	// fmt.Println(repeat(" ", level) + key)
	for k := range data {
		resultBags[k] = true
		search(k, bagGraph, resultBags, level+1)
	}

}

func repeat(s string, times int) string {
	concat := ""
	for i := 0; i < times; i++ {
		concat += s
	}
	return concat
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func trim(data string, patterns []string) string {
	d := data
	for _, pattern := range patterns {
		d = strings.Trim(d, pattern)
	}
	return d
}
