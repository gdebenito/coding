package main

import (
	"fmt"
	"testing"
)

func singleTest(str string, res bool) {
	if ValidBraces(str) != res {
		fmt.Printf("Test failed for %s\n", str)
	}
}

func TestValidBraces(t *testing.T) {
	singleTest("(){}[]", true)
	singleTest("([{}])", true)
	singleTest("(}", false)
	singleTest("[(])", false)
	singleTest("[({)](]", false)
	singleTest("[", false)
	singleTest("[(", false)
	singleTest("[({", false)
	singleTest("}", false)
	singleTest("}]", false)
	singleTest("}])", false)
	singleTest("(})", false)
}
