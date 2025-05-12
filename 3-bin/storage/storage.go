package storage

import (
	"3-bin/bins"
	"3-bin/file"
	"encoding/json"
)

const storageFile = "storage.json"

type Storage interface{
	Save() error
	Load() error
}

type FileStorage struct {
	Storage
	Data bins.BinList
}

func NewStorage(data *bins.BinList) *FileStorage {
	return &FileStorage{
		Data: *data,
	}
}

func (storage *FileStorage) Save() error {
	content, err := json.Marshal(storage)
	if err != nil {
		return err
	}
	err = file.WriteFile(storageFile, string(content))
	return err
}

func (storage *FileStorage) Load() error {
	data, err := file.ReadFile(storageFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &storage)
	return err
}
