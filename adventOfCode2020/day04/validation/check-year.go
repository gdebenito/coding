package validation

import (
	"fmt"
	"strconv"
)

func checkYear(passport map[string]string) bool {
	iyr, ok := passport["iyr"]
	if !ok {
		fmt.Println("-iyr: missing")
		return false
	}

	issueYear, err := strconv.Atoi(iyr)
	if err != nil || issueYear < 2010 || issueYear > 2020 {
		fmt.Println("-iyr: lower than 2010 or greater than 2020")
		return false
	}
	return true
}
