package io

import (
	"io/ioutil"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestNewFileManager(t *testing.T) {
	type args struct {
		file string
		line int
	}
	tests := []struct {
		name    string
		args    args
		want    *FileManager
		wantErr bool
	}{
		{"newFileManager", args{"fixtures/testfile.txt", 0}, &FileManager{file: "fixtures/testfile.txt", line: 0}, false},
		{"newFileManagerWithLine", args{"fixtures/testfile.txt", 10}, &FileManager{file: "fixtures/testfile.txt", line: 10}, false},
		{"newFileManagerWithLine2", args{"testfile.txt", 1}, &FileManager{file: "testfile.txt", line: 1}, false},
		{"newFileManagerNoFile", args{"", 1}, nil, true},
		{"newFileManagerInvalidLine", args{"testfile.txt", -1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFileManager(tt.args.file, tt.args.line)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewFileManager() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileManager_Read(t *testing.T) {
	type fields struct {
		file string
		line int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{"1", fields{file: "fixtures/readfile.txt", line: 0}, []string{"test", "test 2"}, false},
		{"2", fields{file: "testfile.txt", line: 0}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileManager{
				file: tt.fields.file,
				line: tt.fields.line,
			}
			got, err := f.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileManager.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileManager.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileManager_Save(t *testing.T) {
	type fields struct {
		file string
		line int
	}
	type args struct {
		content []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"noError", fields{file: "fixtures/writefile.txt", line: 0}, args{content: []string{"line 1", "line 2"}}, false},
		{"invalidName", fields{file: "test/writefile.txt", line: 0}, args{content: []string{"line 1", "line 2"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileManager{
				file: tt.fields.file,
				line: tt.fields.line,
			}
			if err := f.Save(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("FileManager.Save() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil {
				bytes, _ := ioutil.ReadFile(tt.fields.file)
				have := []byte(strings.Join(tt.args.content, "\n"))

				if !reflect.DeepEqual(bytes, have) {
					t.Errorf("FileManager.Save() content = %v, wantCont %v", bytes, have)
				}
				command := exec.Command("rm", "-f", tt.fields.file)
				command.Run()
			}
		})
	}
}
