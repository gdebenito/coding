package main

import (
	"bufio"
	"fmt"
	"os"
)

// The main strategy to attack this problem is to count the different characters
// the group has, and adding it to a count
func main() {
	fileHandle, _ := os.Open("input.txt")
	// fileHandle, _ := os.Open("test.txt")
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	questionsAnsweredCount := 0
	for scanner.Scan() {
		questionsAnswered := make(map[rune]bool)
		for {
			text := scanner.Text()

			if text == "" {
				// Go to new group
				break
			}

			addQuestionsToGroup(questionsAnswered, text)
			// fmt.Printf("For: "+text+" ; Are: %d\n", len(questionsAnswered))

			if hasNext := scanner.Scan(); !hasNext {
				break
			}
		}
		questionsAnsweredCount += len(questionsAnswered)
	}

	fmt.Printf("Answered questions count: %d\n", questionsAnsweredCount)
}

func addQuestionsToGroup(groupQuestionsAnsweredMap map[rune]bool, questions string) {
	for _, question := range questions {
		groupQuestionsAnsweredMap[question] = true
	}
}
