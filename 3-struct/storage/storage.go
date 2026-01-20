package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"
	"os"
)

const STORAGE_FILE = "storage.json"

func Save(data bins.BinList) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(STORAGE_FILE, jsonData, 0644)
}

func Get() (bins.BinList, error) {
	jsonData, err := file.ReadFile(STORAGE_FILE)

	if err != nil {
		return nil, err
	}

	var data bins.BinList

	err = json.Unmarshal(jsonData, &data)

	return data, err
}
