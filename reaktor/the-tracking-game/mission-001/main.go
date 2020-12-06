package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"os"
)

func main() {
	signal := readFile()
	bytesRead := 0
	mapOfChars := make(map[rune]bool)
	password := ""
	for _, c := range signal {
		char := rune(c)
		if _, isInMap := mapOfChars[char]; isInMap {
			// fmt.Printf("does not match %s with %c\n", password, char)
			password, bytesRead = removeFromRepeatingCharacter(password, char, mapOfChars)
		} else {
			mapOfChars[char] = true
			bytesRead++
			password = password + string(char)
		}

		if bytesRead == 16 {
			d, _ := b64.StdEncoding.DecodeString(password)
			fmt.Println("password is: " + string(d))
		}

	}
}

func removeFromRepeatingCharacter(password string, char rune, mapOfChars map[rune]bool) (string, int) {
	index := 0
	for i, c := range password {
		// When character match with the actual
		if char == rune(c) {
			index = i + 1
			break
		}

		// Delete character from map entry
		delete(mapOfChars, c)
	}

	newPassword := password[index:] + string(char)
	bytesRead := len(newPassword)

	return newPassword, bytesRead

}

func readFile() string {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	scanner.Scan()
	return scanner.Text()
}
