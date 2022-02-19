package gorns

import(
	"testing"
	"github.com/hexacry/gorns/utils"
)

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

func Test_New(t *testing.T) {
	expected := int16(8)

	instance := New()

	got := instance.Cfg.StorageLimit

	if expected != got {
		t.Errorf("\nExpected: %d\nGot: %d\n", expected, got)
	}
}

func Test_Use(t *testing.T) {
	expected := int16(8)

	m := &Middleware{
		Version: make(utils.VersionType, 3),
		Warner: New(),
	}

	got := m.Warner.Cfg.StorageLimit

	if expected != got {
		t.Errorf("\nExpected: %d\nGot: %d\n", expected, got)
	}
}
