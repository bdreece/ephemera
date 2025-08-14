package storage

import (
	"io/fs"
	"os"
)

type RootProvider struct {
	*os.Root
}

// Create implements Provider.
// Subtle: this method shadows the method (*Root).Create of RootProvider.Root.
func (p *RootProvider) Create(name string) (File, error) {
	return p.Root.Create(name)
}

// Mkdir implements Provider.
// Subtle: this method shadows the method (*Root).Mkdir of RootProvider.Root.
func (p *RootProvider) Mkdir(name string, perm fs.FileMode) error {
	return p.Root.Mkdir(name, perm)
}

// Open implements [fs.FS].
// Subtle: this method shadows the method (*Root).Open of RootProvider.Root.
func (p *RootProvider) Open(name string) (fs.File, error) {
	return p.Root.Open(name)
}

// OpenFile implements Provider.
// Subtle: this method shadows the method (*Root).OpenFile of RootProvider.Root.
func (p *RootProvider) OpenFile(name string, flag int, perm fs.FileMode) (File, error) {
	return p.Root.OpenFile(name, flag, perm)
}

// Stat implements [fs.StatFS].
// Subtle: this method shadows the method (*Root).Stat of RootProvider.Root.
func (p *RootProvider) Stat(name string) (fs.FileInfo, error) {
	return p.Root.Stat(name)
}

// Sub implements [fs.SubFS].
func (p *RootProvider) Sub(dir string) (fs.FS, error) {
	root, err := p.OpenRoot(dir)
	if err != nil {
		return nil, err
	}

	return &RootProvider{root}, nil
}
