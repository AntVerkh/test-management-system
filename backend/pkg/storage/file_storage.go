package storage

import (
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FileStorage interface {
	Save(file io.Reader, filename string) (string, error)
	Get(filepath string) (io.ReadCloser, error)
	Delete(filepath string) error
}

type LocalFileStorage struct {
	basePath string
}

func NewLocalFileStorage(basePath string) FileStorage {
	return &LocalFileStorage{basePath: basePath}
}

func (s *LocalFileStorage) Save(file io.Reader, filename string) (string, error) {
	// Create unique filename
	ext := filepath.Ext(filename)
	newFilename := uuid.New().String() + ext
	filePath := filepath.Join(s.basePath, newFilename)

	// Create directory if not exists
	if err := os.MkdirAll(s.basePath, 0755); err != nil {
		return "", err
	}

	// Create file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return filePath, nil
}

func (s *LocalFileStorage) Get(filepath string) (io.ReadCloser, error) {
	return os.Open(filepath)
}

func (s *LocalFileStorage) Delete(filepath string) error {
	return os.Remove(filepath)
}
