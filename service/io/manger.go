package io

import (
	"errors"
	"io/ioutil"
	"strings"
)

// Manager that exposes file operations
type Manager interface {
	Read() ([]string, error)
	Save(content []string) error
}

// FileManager implementation for file manager
type FileManager struct {
	file string
	line int
}

// NewFileManager instances a new FileManager
func NewFileManager(file string, line int) (*FileManager, error) {
	if file == "" {
		return nil, errors.New("File cannot be empty")
	}

	if line < 0 {
		return nil, errors.New("Line cannot be negative")
	}

	return &FileManager{
		file: file,
		line: line,
	}, nil
}

// Read reads the complete file and returns an array of strings that represent each line
func (f *FileManager) Read() ([]string, error) {
	file, err := ioutil.ReadFile(f.file)
	if err != nil {
		return nil, err
	}

	str := string(file)

	arr := strings.Split(str, "\n")

	return arr, nil
}

// Save receives an array of strings and writes them to the file
func (f *FileManager) Save(content []string) error {
	fullContent := strings.Join(content, "\n")
	bytes := []byte(fullContent)

	err := ioutil.WriteFile(f.file, bytes, 0666)
	if err != nil {
		return err
	}

	return nil
}
