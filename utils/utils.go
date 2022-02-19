package utils

import(
	"strings"
	"fmt"
	"strconv"
)

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

type VersionType []uint
type VersionTypeByte []byte

func (v VersionType) ToByte() VersionTypeByte {
	var temp []string// = make([]string, len(v))
	var this VersionTypeByte

	for _, e := range v {
		temp = append(temp, fmt.Sprintf("%d", e))
	}

	this = VersionTypeByte(strings.Join(temp, "."))

	return this
}

func (v VersionTypeByte) ToVersion() VersionType {
	var temp []string = strings.Split(string(v), ".")
	var this VersionType = make(VersionType, len(temp))

	for i, _ := range temp {
		temp2, err := strconv.Atoi(string(temp[i]))
		if err != nil {
			temp2 = 0
		}
		this[i] = uint(temp2)
	}

	return this
}

func Version(v ...uint) VersionType {
	var this VersionType

	for _, e := range v {
		this = append(this, e)
	}

	return this
}
