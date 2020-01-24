package commons

import (
	"strconv"
	"strings"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func LeftPad(s string, padString string, length int) string {
	totalLength := length - len(s)
	if totalLength < 0 {
		return s
	}
	return strings.Repeat(padString, totalLength) + s
}
