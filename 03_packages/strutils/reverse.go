package strutils

import (
	"bytes"
	"strings"
)

// ReverseConcate func
// Need args string
func ReverseConcate(str string) string {
	var index int
	var result string

	split := strings.Split(str, "")
	lens := len(split)

	for i := 0; i < lens; i++ {
		index = (lens - i) - 1
		result = result + split[index]
	}
	return result
}

// ReverseBuffer func
// Need args string
func ReverseBuffer(str string) string {
	var result bytes.Buffer

	split := strings.Split(str, "")
	lens := len(split)

	for i := 0; i < lens; i++ {
		index := (lens - i) - 1
		result.WriteString(split[index])
	}
	return result.String()
}

// ReverseJoin func
// Need args string
func ReverseJoin(str string) string {
	var result []string
	split := strings.Split(str, "")
	lens := len(split)

	for i := 0; i < lens; i++ {
		index := (lens - i) - 1
		result = append(result, split[index])
	}
	return strings.Join(result, "")
}
