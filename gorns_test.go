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

	m := &Manager{
		Version: make(utils.VersionType, 3),
		Warner: New(),
	}

	got := m.Warner.Cfg.StorageLimit

	if expected != got {
		t.Errorf("\nExpected: %d\nGot: %d\n", expected, got)
	}
}

func Test_CreateManager(t *testing.T) {
	type testingMap map[int]interface{}
	var expected testingMap = testingMap{
		1: int16(8),
		2: string(make(utils.VersionType, 3).ToByte()),
		3: "TEST",
	}

	x := CreateManager()

	x.Warner.Storage.Push(&UWarn{
		Name: "TEST",
		Code: utils.Code("TEST"),
		Content: "test",
	})

	got := testingMap{
		1: x.Warner.Cfg.StorageLimit,
		2: string(x.Version.ToByte()),
		3: x.Warner.Storage.Get("TEST").Name,
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != got[i] {
			t.Errorf("\nExpected: %v\nGot: %v\n", expected[i], got[i])
		}
	}
}

func Test_DW(t *testing.T) {
	expected := "DW_WARN_NOT_EXIST"

	if got := New().Get(DW_WNE).Name; expected != got {
		t.Errorf("\nExpected: %s\nGot: %s\n", expected, got)
	}
}

func Test_WarnerPushAndGet(t *testing.T) {
	expected := utils.Code("TEST")
	var got uint16

	warner := New()
	if warn := warner.Push("TEST", "test"); warn != nil { // PUSH TEST
		got = warn.Code
	}

	got = warner.Get("TEST").Code // GET TEST

	if expected != got {
		t.Errorf("\nExpected: %d\nGot: %d\n", expected, got)
	}
}

func Test_WarnerUpdate(t *testing.T) {
	expected := "TEST_UPDATED"

	warner := New()
	_ = warner.Push("TEST", "test")
	_ = warner.Update("TEST", &UWarn{
		Name: "TEST_UPDATED",
		Code: utils.Code("TEST_UPDATED"),
		Content: "test updated",
	})

	got := warner.Get("TEST").Name

	if expected != got {
		t.Errorf("\nExpected: %s\nGot: %s\n", expected, got)
	}
}
