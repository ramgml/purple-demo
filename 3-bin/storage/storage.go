package storage

import (
	"3-bin/bins"
	"3-bin/file"
	"encoding/json"
)

type Storage interface{
	Save() error
	Load() error
}

type FileStorage struct {
	Data bins.BinList
}

func NewStorage(data *bins.BinList) Storage {
	return &FileStorage{
		Data: *data,
	}
}

func (storage *FileStorage) Save() error {
	content, err := json.Marshal(storage)
	if err != nil {
		return err
	}
	err = file.WriteFile("storage.json", string(content))
	return err
}

func (storage *FileStorage) Load() error {
	data, err := file.ReadFile("storage.json")
	if err != nil {
		return err
	}
	var tmp *FileStorage
	err = json.Unmarshal(data, &tmp)
	if err == nil {
		storage = tmp
	}
	return err
}
