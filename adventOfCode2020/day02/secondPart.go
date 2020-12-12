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
		index1, err := strconv.Atoi(minAndMax[0])
		if err != nil {
			panic(err)
		}
		index2, err := strconv.Atoi(minAndMax[1])
		if err != nil {
			panic(err)
		}

		letter := parts[1][0]
		password := parts[2]

		var exactlyOne = false

		if len(password) >= index1 && letter == password[index1-1] {
			exactlyOne = !exactlyOne

		}

		if len(password) >= index2 && letter == password[index2-1] {
			exactlyOne = !exactlyOne

		}

		if exactlyOne {
			numberOfValidPasswords++
			// fmt.Printf("Valid %s\n", line)
			// } else {
			// 	fmt.Printf("inValid %s\n", line)
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", numberOfValidPasswords)
	fmt.Printf("Count: %d\n", count)

}
