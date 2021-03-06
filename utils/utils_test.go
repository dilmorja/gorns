package utils

import "testing"

func Test_CharCodeAt(t *testing.T) {
	expected := 84

	got := int(CharCodeAt("T", 0))

	if expected != got {
		t.Errorf("Expected: %d, got: %d", expected, got)
	}
}

func Test_Code(t *testing.T) {
	expected := 84

	c := Code("TEST")

	got := int(c)

	if expected != got {
		t.Errorf("Expected: %d, got: %d", expected, got)
	}
}

func Test_ToByte(t *testing.T) {
	expected := []byte("0.0.0")

	p := make(VersionType, 3)

	got := p.ToByte()

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("\nExpected: %d\nGot: %d\n", e, got[i])
		}
	}
}

func Test_ToVersion(t *testing.T) {
	expected := make(VersionType, 3)

	x := make(VersionType, 3)
	y := x.ToByte()

	got := y.ToVersion()

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("\nExpected: %d\nGot: %d\n", e, got[i])
		}
	}
}

func Test_Version(t *testing.T) {
	expected := make(VersionType, 3)

	got := Version(0,0,0)

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("\nExpected: %d\nGot: %d\n", e, got[i])
		}
	}
}
