package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Archiver is an interface for different archiving types
type Archiver interface {
	DestFmt() string
	Archive(src, dest string) error
}

type zipper struct{}

// ZIP is an Archiver that zips and unzips files.
var ZIP Archiver = (*zipper)(nil)

func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil // skip directories
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		f, err := w.Create(path)
		if err != nil {
			return err
		}

		io.Copy(f, in)
		return nil
	})
}

func (z *zipper) DestFmt() string {
	return "%d.zip"
}
