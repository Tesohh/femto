package storage

import (
	"encoding/json"
	"log/slog"
	"os"
	"path"
)

type PersistentStorage[T any] struct {
	Id   string
	Data T
}

func (p PersistentStorage[T]) Load() error {
	file, err := os.Open(path.Join(".data", p.Id+".json"))
	if err != nil {
		return err
	}
	defer file.Close()
	var data T
	return json.NewDecoder(file).Decode(data)
}

func (p PersistentStorage[T]) Save() error {
	file, err := os.Open(path.Join(".data", p.Id+".json"))
	slog.Info(file.Name())
	if err != nil {
		return err
	}
	defer file.Close()
	var data T
	return json.NewEncoder(file).Encode(data)
}
