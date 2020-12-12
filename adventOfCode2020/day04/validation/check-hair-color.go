package validation

func checkHairColor(passport map[string]string) bool {
	charactersAccepted := map[rune]bool{
		'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true,
		'0': true, '1': true, '2': true, '3': true, '4': true, '5': true,
		'6': true, '7': true, '8': true, '9': true,
	}
	hcl, ok := passport["hcl"]
	if !ok {
		return false
	}
	hairColor := hcl

	if hairColor[0] != '#' {
		return false
	}

	if len(hairColor[1:]) != 6 {
		return false
	}

	for _, character := range hairColor[1:] {
		if _, ok := charactersAccepted[character]; !ok {
			return false
		}
	}

	return true
}
