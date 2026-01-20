package file

import (
	"errors"
	"os"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	fileInfo, err := os.Stat(filename)

	if err != nil {
		return nil, errors.New("Ошибка при получении информации о файле")
	}

	if !strings.HasSuffix(fileInfo.Name(), ".json") {
		return nil, errors.New("Файл должен быть в формате JSON")
	}

	return os.ReadFile(filename)
}
