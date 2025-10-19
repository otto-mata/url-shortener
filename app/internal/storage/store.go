package storage

// Store is the storage abstraction for short links
 type Store interface {
	Save(code, url string) error
	Get(code string) (string, bool)
 }
