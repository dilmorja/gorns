package gorns

import "testing"

func Test_UWarn(t *testing.T) {
	expected := "TEST"

	var x *UWarn = &UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	}

	got := x.Name

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func Test_Swarnf(t *testing.T) {
	expected := "84"

	var x *UWarn = &UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	}

	got := x.Swarnf("%d", x.Code)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}
