package utils

import "strings"

// Get the character code in the indicated position.
func CharCodeAt(str string, at int) rune {
	return []rune(str)[at]
}

// Generates a code based on a text string.
// It is recommended that the text string be in uppercase
// and words should be separated by underscores.
// This is used for the structure of UWarn
//		uwarn := &UWarn{
//			Name: "TEST",
//			Code: utils.Code("TEST"), // return 84
//			Content: "test",
//		}
func Code(str string) uint16 {
	var code uint16

	a := strings.Split(str, "_")

	for i := range a {
		if code == 0xFDE8 { // An approximation to the uint16 built-in type range
			break
		}
		code = (code + uint16(CharCodeAt(a[i], 0)))
	}

	return code

}
