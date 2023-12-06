package chars

func CharInRange(char rune, start rune, end rune) bool {
	return char >= start && char <= end
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
