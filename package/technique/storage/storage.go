package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type LocalFileStorage struct {
	basePath string
}

func (s *LocalFileStorage) SaveFile(filename string, data []byte) error {
	filepath := filepath.Join(s.basePath, filename)
	return ioutil.WriteFile(filepath, data, 0644)
}

func (s *LocalFileStorage) GetFile(filename string) ([]byte, error) {
	filepath := filepath.Join(s.basePath, filename)
	return ioutil.ReadFile(filepath)
}

func (s *LocalFileStorage) DeleteFile(filename string) error {
	filepath := filepath.Join(s.basePath, filename)
	return os.Remove(filepath)
}
