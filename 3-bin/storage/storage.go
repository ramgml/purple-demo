package storage

import (
	"3-bin/bins"
	"3-bin/file"
	"encoding/json"
)

// Работа с файлом
type Storage struct {
	Data bins.BinList
}

func NewStorage(data *bins.BinList) Storage {
	var dataStorage Storage
	dataStorage.Data = *data
	return dataStorage
}

func (storage *Storage) Save() error {
	content, err := json.Marshal(storage)
	if err != nil {
		return err
	}
	err = file.WriteFile("storage.json", string(content))
	return err
}

func (storage *Storage) List() (*Storage, error) {
	data, err := file.ReadFile("storage.json")
	if err != nil {
		return storage, err
	}
	var tmp *Storage
	err = json.Unmarshal(data, &tmp)
	if err == nil {
		storage = tmp
	}
	return storage, err
}
