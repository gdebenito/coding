package main

var m = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func ValidBraces(str string) bool {
	s := []rune{}

	if len(str)%2 != 0 {
		return false
	}

	for _, v := range str {
		switch v {
		case '{', '[', '(':
			s = append(s, v)
		case '}', ']', ')':
			if len(s) > 0 {
				lastBrace := s[len(s)-1]
				if m[lastBrace] == v {
					// Pop last element
					s = s[:len(s)-1]
				} else {
					// Last open brace doesn't match current closing brace
					return false
				}

			} else {
				return false
			}

		default:
			// There aren't more cases
			return false
		}
	}

	// Remaining elements in the s
	return len(s) == 0
}
