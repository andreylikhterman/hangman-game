package slice

func Contains(array []rune, element rune) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}

	return false
}
