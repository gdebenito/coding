package validation

import "strconv"

func checkExpirationYear(passport map[string]string) bool {
	eyr, ok := passport["eyr"]
	if !ok {
		return false
	}
	expirationYear, err := strconv.Atoi(eyr)
	if err != nil || expirationYear < 2020 || expirationYear > 2030 {
		return false
	}
	return true
}
