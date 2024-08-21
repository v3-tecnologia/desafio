package inmemory_adapter

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

func NewStorageInMemoryAdapter() *StorageInMemoryAdapter {
	return &StorageInMemoryAdapter{
		Files: make(map[string]*bytes.Buffer),
	}
}

type StorageInMemoryAdapter struct {
	Files map[string]*bytes.Buffer
	mu    sync.Mutex
}

func (s *StorageInMemoryAdapter) SaveFile(file io.Reader, fileName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.Files[fileName]; exists {
		return errors.New("file already exists")
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return err
	}

	s.Files[fileName] = &buf
	return nil
}
