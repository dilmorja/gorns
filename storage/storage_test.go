package storage

import (
	"reflect"
	"testing"

	"github.com/pszao/gorns"
	"github.com/pszao/gorns/utils"
)

func Test_Storage(t *testing.T) {
	expected := "*storage.Storage"

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

	se.Push(&gorns.UWarn{
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

	_ = se.Push(&gorns.UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	})

	warn := se.Push(&gorns.UWarn{
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

	se.Push(&gorns.UWarn{
		Name:    "TEST",
		Code:    uint16(84),
		Content: "test",
	})

	got := se.Push(&gorns.UWarn{
		Name:    "TEST_TWO",
		Code:    uint16(168),
		Content: "test",
	})

	if got != nil {
		t.Errorf("Expected: %v, got: %v", nil, got)
	}
}

func Tets_Delete(t *testing.T) {
	expected := 0

	storage := NewStorage(&StorageConfig{
		Limit: int16(6),
	})

	warn := storage.Push(&gorns.UWarn{
		Name:    "TEST",
		Code:    utils.Code("TEST"),
		Content: "test",
	})

	if warn != nil {
		t.Errorf("warn: %v", warn)
	}

	got := len(storage.warns)

	if expected != got {
		t.Errorf("Expected: %d, got: %d", expected, got)
	}
}