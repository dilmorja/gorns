package storage

import (
	"sync"

	"github.com/pszao/gorns"
)

// This structure is private to prevent inappropriate or unwanted use.
type storage struct {
	warns map[string]*gorns.UWarn
	mutex sync.RWMutex
	limit int16
}

// The storage structure is like the cache, only its elements do not expire.
// Instead, you can and should set a space limit for storing new items.
type Storage struct {
	*storage
}

// The configuration structure required to create a new storage instance.
type StorageConfig struct {
	// This limit must be carefully set.
	// If incorrectly assigned,
	// it could result in an unexpected crash of the application,
	// or otherwise uninterrupted execution (unless there is some fatal runtime error).
	Limit int16
}

// Create a new storage instance.
//		storage := gorns.NewStorage(&StorageConfig{
//			Limit: int16(6), // Converting from int to int16 is recommended.
//		})
// If the storage variable is incorrectly modified,
// there is a risk of losing the warnings stored so far.
func NewStorage(sc *StorageConfig) Storage {

	this := new(storage)
	this.limit = sc.Limit
	this.warns = make(map[string]*gorns.UWarn)
	this.mutex = sync.RWMutex{}

	return Storage{
		this,
	}

}

// Save a new warning if it doesn't exist;
// otherwise it returns the existing warning.
// If the warning does not exist and the storage limit has not been reached;
// continue with the build process; otherwise return nil.
//		// ...
//		warn := storage.Push(&UWarn{
//			Name: "EXAMPLE",
//			Code: uint16(69),
//			Content: "Hi! This is a example. :D",
//		})
//		// ...
func (s *storage) Push(warn *gorns.UWarn) *gorns.UWarn {
	if w := s.Get(warn.Name); w != nil {
		return w
	}
	if int16(len(s.warns)) != s.limit {
		s.mutex.Lock()
		s.warns[warn.Name] = warn
		s.mutex.Unlock()
	}
	return nil
}

// Get the object of a warning if it exists; otherwise it returns nil.
//		// ...
//		getWarn := storage.Get("EXAMPLE")
//		// ...
func (s *storage) Get(name string) *gorns.UWarn {
	s.mutex.RLock()
	if w, ok := s.warns[name]; ok {
		s.mutex.RUnlock()
		return w
	}
	s.mutex.RUnlock()
	return nil
}

// Removes a warning from storage and returns true if the operation was successful.
//		// ...
//		storage.Delete("TEST")
//		// ...
func (s *storage) Delete(name string) bool {
	s.mutex.RLock()
	if _, ok := s.warns[name]; ok {
		s.mutex.RUnlock()
		s.mutex.Lock()
		delete(s.warns, name)
		s.mutex.Unlock()
		return true
	}
	return false
}
