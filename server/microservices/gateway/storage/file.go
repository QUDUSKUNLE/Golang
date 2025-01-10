package storage

import (
	"bytes"
)

type File struct {
	name   string
	buffer *bytes.Buffer
}

func NewFile(name string) *File {
	return &File{
		name:   name,
		buffer: new(bytes.Buffer),
	}
}

func (f *File) Write(p []byte) error {
	_, err := f.buffer.Write(p)
	return err
}
