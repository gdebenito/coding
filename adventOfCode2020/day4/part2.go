package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./validation"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	numberOfPassportsValid := 0
	for scanner.Scan() {
		passport := make(map[string]string)
		for {
			line := scanner.Text()

			// Empty line must check the data collected and start new passport
			if line == "" {
				printMap(passport)
				if validation.ValidatePassport(passport) {
					numberOfPassportsValid++
				}
				break
			}

			values := strings.Split(line, " ")

			fillPassportWithData(passport, values)

			hasNext := scanner.Scan()
			if !hasNext {
				break
			}
		}
	}

	fmt.Printf("Number of passports valid: %d", numberOfPassportsValid)
}

func fillPassportWithData(passport map[string]string, values []string) {
	for _, value := range values {
		keyAndValue := strings.Split(value, ":")
		// fmt.Println(keyAndValue)
		key := keyAndValue[0]
		value := keyAndValue[1]
		passport[key] = value
	}
}

// https://stackoverflow.com/questions/27117896/how-to-pretty-print-variables
func printMap(m map[string]string) {
	var maxLenKey int
	for k := range m {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
	}

	for k, v := range m {
		fmt.Printf(k + ":" + strings.Repeat(" ", maxLenKey-len(k)) + v + " ")
	}

	fmt.Printf("\n")
}
