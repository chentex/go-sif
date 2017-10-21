package service

import (
	"github.com/chentex/go-sif/service/io"
	"github.com/pkg/errors"
)

//InsertServiceInterface defines the methods that any inserter need to follow
type InsertServiceInterface interface {
	Insert(file string, line int, text string) error
	SetManager(m io.Manager)
}

//InsertService Implementing a file based insert service
type InsertService struct {
	fm io.Manager
}

//NewInsertService return a new instance of InsertService
func NewInsertService() InsertServiceInterface {
	return &InsertService{}
}

//SetManager assigns a new file manager
func (i *InsertService) SetManager(m io.Manager) {
	i.fm = m
}

//Insert text into file in given line
func (i *InsertService) Insert(file string, line int, text string) error {
	if file == "" {
		return errors.Wrap(errors.New("file parameter cannot be empty"), "cmd.Insert validations")
	}

	if line < -1 || line == 0 {
		return errors.Wrap(errors.New("line parameter with invalid value. Must be 1 and above or not set"), "cmd.Insert validations")
	}

	content, err := i.fm.Read(file)
	if err != nil {
		return errors.Wrap(err, "service.InsertService reading file")
	}

	len := len(content)
	if line > len {
		return errors.Wrap(errors.New("line parameter out of bounds"), "cmd.Insert validations")
	}

	if line == -1 {
		line = len + 1
	}

	preContent := make([]string, (line - 1))
	copy(preContent, content[:line-1])
	postContent := content[line-1:]

	finalContent := append(preContent, text)
	finalContent = append(finalContent, postContent...)

	err = i.fm.Save(finalContent, file)
	if err != nil {
		return errors.Wrap(err, "service.InsertService writing file")
	}

	return nil
}
