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
		lineNumber := 0
		for {
			text := scanner.Text()

			if text == "" {
				// Go to new group
				break
			}

			if lineNumber == 0 {
				addQuestionsToGroup(questionsAnswered, text)
			} else {
				questionMap := make(map[rune]bool)
				addQuestionsToGroup(questionMap, text)
				intersectAnswers(questionsAnswered, questionMap)
			}

			lineNumber++

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

func intersectAnswers(groupQuestionsAnsweredMap map[rune]bool, questionsNew map[rune]bool) {
	for key := range groupQuestionsAnsweredMap {
		if _, isInMap := questionsNew[key]; !isInMap {
			delete(groupQuestionsAnsweredMap, key)
		}
	}
}
