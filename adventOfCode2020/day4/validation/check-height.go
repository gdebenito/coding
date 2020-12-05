package validation

import (
	"fmt"
	"strconv"
)

func checkHeight(passport map[string]string) bool {
	hgt, ok := passport["hgt"]

	if !ok {
		fmt.Println("-hgt: missing")
		return false
	}
	inchesOrCm := hgt[len(hgt)-2:]
	height, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		fmt.Println("-hgt: non digit")
		return false
	}

	if inchesOrCm != "in" && inchesOrCm != "cm" {
		fmt.Println("-hgt: neither inches or cm")
		return false
	}

	if inchesOrCm == "in" && (height < 59 || height > 76) {
		fmt.Println("-hgt: inches lower than 59 or greater than 76")
		return false
	}

	if inchesOrCm == "cm" && (height < 150 || height > 193) {
		fmt.Println("-hgt: cm lower than 150 or greater than 193")
		return false
	}
	return true

}
