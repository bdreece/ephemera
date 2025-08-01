package storage

import (
	"io"
	"io/fs"
)

type File interface {
	fs.ReadDirFile
	io.ReaderAt
	io.WriterAt
	io.WriteSeeker
}

type Provider interface {
	fs.SubFS
	fs.StatFS

	Name() string
	Create(name string) (File, error)
	Mkdir(name string, perm fs.FileMode) error
	OpenFile(name string, flag int, perm fs.FileMode) (File, error)
	Remove(name string) error
}
