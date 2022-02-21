package gorns

import(
	"fmt"
	"github.com/hexacry/gorns/utils"
)

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
	Swarnf() string
}

// Formats an element of type UWarn and returns a string.
// Its use is perfect for terminals, web applications, etc.
func (uw *UWarn) Swarnf(format string, v ...interface{}) string {
	if len(v) > 0 {
		return fmt.Sprintf(format, v...)
	}

	return fmt.Sprintf("%s (%d): %s", uw.Name, uw.Code, uw.Content)
}


type Warner struct {
	Storage Storage
	Cfg *WarnerOpts
}

func (w *Warner) Push(name string, content string) *UWarn {
	return w.Storage.Push(&UWarn{
		Name: name,
		Code: utils.Code(name),
		Content: content,
	})
}

func (w *Warner) Get(name string) *UWarn {
	return w.Storage.Get(name)
}

func (w *Warner) Delete(name string) bool {
	return w.Storage.Delete(name)
}

func (w *Warner) Update(name string, new *UWarn) bool {
	return w.Storage.Update(name, new)
}

func (w *Warner) Swarnf(name string) string {
	return w.Get(name).Swarnf()
}

type WarnerOpts struct {
	StorageLimit int16
}

func New(opts ...*WarnerOpts) *Warner {
	var this *Warner = new(Warner)

	if len(opts) > 0 {
		this.Cfg = opts[0]
	}

	if this.Cfg == nil {
		this.Cfg = &WarnerOpts{
			StorageLimit: int16(8),
		}
	}

	this.Storage = NewStorage(&StorageConfig{
		Limit: this.Cfg.StorageLimit,
	})

	return this
}

type Manager struct {
	Version utils.VersionType
	Warner *Warner
}

func Use(m Manager) *Warner {
	return m.Warner
}

func CreateManager() *Manager {
	var this *Manager = &Manager{
		Version: utils.Version(0,0,0),
		Warner: New(),
	}
	return this
}
