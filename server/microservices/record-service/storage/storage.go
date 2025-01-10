package storage

import (
	"os"
)

type Manager interface {
	Store(file *File) error
}

type Storage struct {
	dir string
}

// Store implements Manager.
func (s Storage) Store(file *File) error {
	err := os.WriteFile(s.dir+file.name, file.buffer.Bytes(), 0644);
	if err != nil {
		return err
	}
	return nil
}

var _ Manager = &Storage{}
