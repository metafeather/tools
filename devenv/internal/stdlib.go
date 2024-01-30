package internal

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/metafeather/tools/devenv/internal/log"
)

var (
	ErrStdlibExists    = errors.New("Stdlib already exists")
	ErrStdlibNotExists = errors.New("Stdlib does not exist")
)

type Stdlib struct {
	efs embed.FS
}

// NewStdlib creates a new Stdlib instance with the given embedded filesystem.
func NewStdlib(efs embed.FS) *Stdlib {
	return &Stdlib{
		efs: efs,
	}
}

// Read attempts to open and read the stdlib directory at the given path.
// It returns ErrStdlibNotExists if the path does not exist, ErrStdlibExists if
// the path already contains a stdlib, or another error.
func (sl *Stdlib) Read(p string) error {
	f, err := os.Open(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("stdlib: %w", ErrStdlibNotExists)
		}
		return err
	}
	defer f.Close()
	return fmt.Errorf("stdlib: %w", ErrStdlibExists)
}

// Write the embedded filesystem to the target directory for use by tools.
// It walks the filesystem, creates any necessary directories,
// writes each file to the corresponding path in the output directory.
func (sl *Stdlib) Write(p string) error {
	files, err := walkFS(sl.efs)
	if err != nil {
		return err
	}

	for _, file := range files {
		content, err := sl.efs.ReadFile(file)
		if err != nil {
			return err
		}
		fullpath := path.Join(p, file)
		err = os.MkdirAll(filepath.Dir(fullpath), 0700)
		if err != nil {
			return err
		}
		if err := os.WriteFile(fullpath, content, 0644); err != nil {
			log.Error("error os.WriteFile error: %v", err)
			return err
		}
		log.Debug("Stdlib written:", "file", file)
	}
	return nil
}

// walkFS walks the given embedded filesystem recursively, and returns a
// slice of all the file paths found.
func walkFS(efs embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, p)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
