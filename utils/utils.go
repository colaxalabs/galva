package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// IsQuoted returns true/false if input is quoted
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

// ParseAddress returns Ethereum address as string
func ParseAddress(input string) string {
	return common.HexToAddress(input).String()
}

// Format Unix time to time.RFC3339
func ParseTime(unixTime int64) (time.Time, error) {
	unixFormat := time.Unix(unixTime, 0).Format(time.RFC3339)
	newTime, _ := time.Parse(time.RFC3339, unixFormat)
	return newTime, nil
}

// HasHexPrefix returns true if input starts with 0x
func HasHexPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}
