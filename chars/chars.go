package chars

func CharInRange(char rune, start rune, end rune) bool {
	if start <= end {
		return char >= start && char <= end
	} else {
		return char >= end && char <= start
	}
}

func CharIsDigit(char rune) bool {
	return CharInRange(char, '0', '9')
}

func CharIsLower(char rune) bool {
	return CharInRange(char, 'a', 'z')
}

func CharIsUpper(char rune) bool {
	return CharInRange(char, 'A', 'Z')
}
