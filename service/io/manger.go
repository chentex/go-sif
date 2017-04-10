package io

import (
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// Manager that exposes file operations
type Manager interface {
	Read(file string) ([]string, error)
	Save(content []string, file string) error
}

// fileManager implementation for file manager
type fileManager struct{}

// NewFileManager instances a new FileManager
func NewFileManager() Manager {
	return &fileManager{}
}

// Read reads the complete file and returns an array of strings that represent each line
func (f *fileManager) Read(file string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "fileManager while reading file")
	}

	str := string(fileBytes)

	arr := strings.Split(str, "\n")

	return arr, nil
}

// Save receives an array of strings and writes them to the file
func (f *fileManager) Save(content []string, file string) error {
	fullContent := strings.Join(content, "\n")
	bytes := []byte(fullContent)

	err := ioutil.WriteFile(file, bytes, 0666)
	if err != nil {
		return errors.Wrap(err, "fileManager while writing file")
	}

	return nil
}
