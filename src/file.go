package src

import (
	"os"
)

type File struct {
	D    []byte
	Size int64 // bytes
}

func new_file(d []byte, size int64) *File {
	return &File{
		D:    d,
		Size: size,
	}
}

func ReadFile(path string) (*File, error) {
	d, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return new_file(d, f.Size()), nil
}
