package validation

import (
	"fmt"
	"strconv"
)

func checkPassportID(passport map[string]string) bool {
	i, ok := passport["pid"]
	if !ok {
		fmt.Println("-pid: Missing")
		return false
	}
	if len(i) != 9 {
		fmt.Println("-pid: length != 9")
		return false
	}

	if _, err := strconv.Atoi(i); err != nil {
		fmt.Println("-pid: non digit characters")
		return false
	}

	return true
}
