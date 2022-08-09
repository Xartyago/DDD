package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type FileStore struct {
	FilePath string
}

func NewStore(fileName string) Store {
	return &FileStore{
		FilePath: fileName,
	}
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}
func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
