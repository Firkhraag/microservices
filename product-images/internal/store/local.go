package store

import (
	"io"
	"path/filepath"
)

// Local is an implementation of the Storage interface
// works with the local disk on the current machine
type Local struct {
	maxFileSize int
	basePath    string
}

// NewLocal creates a new local files store
// basePath is the base directory to save files to
// maxSize is the max number of bytes that a file can be
func NewLocal(maxFileSize int, basePath string) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{maxFileSize, p}, nil
}

// Save the contents to the given path
// path is a relative path, basePath will be appended
func (l *Local) Save(path string, contents io.Reader) error {
	return nil
}
