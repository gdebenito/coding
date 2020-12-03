package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	numberOfValidPasswords := 0
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		count++
		line := scanner.Text()
		parts := strings.Split(line, " ")
		minAndMax := strings.Split(parts[0], "-")
		min, err := strconv.Atoi(minAndMax[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(minAndMax[1])
		if err != nil {
			panic(err)
		}

		letter := rune(parts[1][0])
		password := parts[2]

		counter := 0
		for _, l := range password {
			if l == letter {
				counter++
			}
		}

		if min <= counter && counter <= max {
			numberOfValidPasswords++
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", numberOfValidPasswords)
	fmt.Printf("Count: %d\n", count)

}
