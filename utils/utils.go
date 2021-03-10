package utils

func IsQuoted(input []byte) bool {
	return len(input) >= 2 &&
		((input[0] == '"' && input[len(input)-1] == '"') ||
			(input[0] == '\'' && input[len(input)-1] == '\''))
}

// RemoveQuotes removes the first and last if they are both ' or ", otherwise its a noop
func RemoveQuotes(input []byte) []byte {
	if IsQuoted(input) {
		return input[1 : len(input)-1]
	}
	return input
}
