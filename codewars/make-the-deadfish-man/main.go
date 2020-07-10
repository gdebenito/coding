package kata

func Parse(data string) []int {
	// Declare result array
	res := []int{}
	size := len(data)
	i := 0
	value := 0

	for i < size {
		switch letter := data[i]; letter {
		case 'i':
			value++
		case 'd':
			value--
		case 'o':
			res = append(res, value)
		case 's':
			value = value * value
		}
		i++
	}

	return res
}
