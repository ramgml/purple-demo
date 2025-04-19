package file

import (
	"os"
	"strings"
	"errors"
)

func WriteFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

func ReadFile(path string) ([]byte, error) {
	if !strings.HasSuffix(path, ".json") {
		return nil, errors.New("IT'S NOT JSON-FILE")
	}
	data, err := os.ReadFile(path)
	return data, err
}
