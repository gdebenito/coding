package validation

import "strconv"

func checkBirthday(passport map[string]string) bool {
	byr, ok := passport["byr"]

	if !ok {
		return false
	}

	birthday, err := strconv.Atoi(byr)
	if err != nil || birthday < 1920 || birthday > 2002 {
		return false
	}

	return true
}
