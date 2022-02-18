package gorns

import (
	"reflect"
	"testing"

	"github.com/hexacry/gorns/utils"
)

func Test_Storage(t *testing.T) {
	expected := "*gorns.Storage"

	st := new(Storage)

	got := reflect.TypeOf(st).String()

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func Test_Push(t *testing.T) {
	expected := "TEST"

	se := NewStorage(&StorageConfig{
		Limit: int16(6),
	})

	se.Push(&UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	})

	got := se.Get("TEST").Name

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func Test_Push_Same(t *testing.T) {
	expected := "TEST"
	var got string

	se := NewStorage(&StorageConfig{
		Limit: int16(6),
	})

	_ = se.Push(&UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	})

	warn := se.Push(&UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	})

	if warn != nil {
		got = warn.Name
	} else {
		got = se.Get("TEST").Name
	}

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func Test_PushLimit(t *testing.T) {

	se := NewStorage(&StorageConfig{
		Limit: int16(1),
	})

	se.Push(&UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	})

	got := se.Push(&UWarn{
		Name:    "TEST_TWO",
		Code:    uint16(168),
		Content: "test",
	})

	if got != nil {
		t.Errorf("Expected: %v, got: %v", nil, got)
	}
}

func Test_Delete(t *testing.T) {
	expected := 0

	storage := NewStorage(&StorageConfig{
		Limit: int16(6),
	})

	warn := storage.Push(&UWarn{
		Name:    "TEST",
		Code:    utils.Code("TEST"),
		Content: "test",
	})

	if warn != nil {
		t.Errorf("warn: %v", warn)
	}

	deleted := storage.Delete("TEST")
	if !deleted {
		t.Errorf("\nTEST not deleted\n")
	}

	got := len(storage.warns)

	if expected != got {
		t.Errorf("Expected: %d, got: %d", expected, got)
	}
}

func Test_Update(t *testing.T) {
	expected := "UPDATED_TEST"

	storage := NewStorage(&StorageConfig{
		Limit: int16(6),
	})

	if warn := storage.Push(&UWarn{
		Name:    "TEST",
		Code:    utils.Code("TEST"),
		Content: "test",
	}); warn != nil {
		t.Errorf("warn: %v", warn)
	}

	if ok := storage.Update("TEST", &UWarn{
		Name:    "UPDATED_TEST",
		Code: 	 utils.Code("UPDATED_TEST"),
		Content: "A updated test",
	}); !ok {
		t.Errorf("ok: %v", ok)
	}

	got := storage.warns["TEST"].Name

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}
