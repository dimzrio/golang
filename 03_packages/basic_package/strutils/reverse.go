package strutils

import (
	"bytes"
	"strings"
)

/*
Note:
Func name with Capitalize is Public --> accessable from another package
Func name with Lowercase is Private --> can't accessable from another package
*/

// Private
func info() string {
	return "This is an strutils package"
}

//GetInfo fuct
func GetInfo() string {
	return info()
}

// ReverseConcate func
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
