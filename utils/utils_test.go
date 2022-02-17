package utils

import (
	"testing"

	"github.com/hexacry/gorns"
)

func Test_CharCodeAt(t *testing.T) {
	expected := 84

	got := int(CharCodeAt("T", 0))

	if expected != got {
		t.Errorf("Expected: %d, got: %d", expected, got)
	}
}

func Test_Code(t *testing.T) {
	expected := 84

	uwarn := &gorns.UWarn{
		Name:    "TEST",
		Code:    Code("TEST"),
		Content: "test",
	}

	got := int(uwarn.Code)

	if expected != got {
		t.Errorf("Expected: %d, got: %d", expected, got)
	}
}
