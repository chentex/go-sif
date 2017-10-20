package io

import (
	"io/ioutil"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestNewFileManager(t *testing.T) {
	tests := []struct {
		name string
		want *fileManager
	}{
		{"newFileManager", &fileManager{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFileManager()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileManager_Read(t *testing.T) {
	type fields struct {
		file string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{"1", fields{file: "fixtures/readfile.txt"}, []string{"test", "test 2"}, false},
		{"2", fields{file: "testfile.txt"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileManager{}
			got, err := f.Read(tt.fields.file)
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
			f := &fileManager{}
			if err := f.Save(tt.args.content, tt.fields.file); (err != nil) != tt.wantErr {
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
