package gorns

import "fmt"

// UWarn (Unwrapped Warn) is the key structure for creating a new warning.
// It is recommended to use it as a pointer.
//	x := &UWarn{
//		Name: "SIMPLE_UWARN",
//		Code: uint16(168),
//		Content: "This is a simple UWarn"
//	}
type UWarn struct {
	// The name of UWarn type element
	Name string
	// This Code is not random or by choice,
	// it has its use and its way of generating it
	Code uint16
	// The UWarn content
	Content string
}

// This should be used to contain warning information.
type Warn interface {
	// Warn type as string
	Wrap() string
	// Wrapped UWarn type to Unwrapped UWarn type
	Unwrap() *UWarn
}

// Formats an element of type UWarn and returns a string.
// Its use is perfect for terminals, web applications, etc.
func (uw *UWarn) Wrap(format string, v ...interface{}) string {
	if len(v) <= 2 {
		return fmt.Sprintf(format, v...)
	}

	return fmt.Sprintf("%s (%d): %s", uw.Name, uw.Code, uw.Content)
}

// Returns an element of type UWarn that was previously wrapped
func (uw *UWarn) Unwrap() *UWarn {
	return uw
}
