package adapter

import (
	"io"
	"os"
	"path/filepath"
)

func NewLocalStorageAdapter(directory string) *LocalStorageAdapter {
	return &LocalStorageAdapter{
		directory: directory,
	}
}

type LocalStorageAdapter struct {
	directory string
}

func (s *LocalStorageAdapter) SaveFile(file io.Reader, fileName string) error {
	if _, err := os.Stat(s.directory); os.IsNotExist(err) {
		err := os.MkdirAll(s.directory, 0755)
		if err != nil {
			return err
		}
	}

	filePath := filepath.Join(s.directory, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
