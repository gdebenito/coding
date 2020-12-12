package validation

func checkEyeColor(passport map[string]string) bool {
	eyeColor, ok := passport["ecl"]
	if !ok {
		return false
	}
	colors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if _, ok := colors[eyeColor]; !ok {
		return false
	}

	return true
}
