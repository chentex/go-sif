package service

import (
	"os"
	"reflect"
	"testing"

	"github.com/chentex/go-sif/service/io"
	"github.com/pkg/errors"
)

func TestNewInsertService(t *testing.T) {
	tests := []struct {
		name string
		want InsertServiceInterface
	}{
		{"#1", NewInsertService()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInsertService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInsertService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertService_SetManager(t *testing.T) {
	type args struct {
		m io.Manager
	}
	tests := []struct {
		name string
		args args
	}{
		{"#1", args{io.NewFileManager()}},
		{"#2", args{newFileManager()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InsertService{}
			i.SetManager(tt.args.m)
		})
	}
}

func TestInsertService_Insert(t *testing.T) {
	i := &InsertService{}
	i.SetManager(newFileManager())
	type args struct {
		file string
		line int
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"#1", args{"file", -2, "<--insert text-->"}, true},
		{"#2", args{"file", 0, "<--insert text-->"}, true},
		{"#3", args{"file", 4, "<--insert text-->"}, true},
		{"#4", args{"file", 1, "<--insert text-->"}, false},
		{"#5", args{"file", 2, "<--insert text-->"}, false},
		{"#6", args{"file", 3, "<--insert text-->"}, true},
		{"#7", args{"", 1, "<--insert text-->"}, true},
		{"#8", args{"file2", 1, "<--insert text-->"}, true},
		{"#9", args{"file3", 1, "<--insert text-->"}, true},
		{"#10", args{"file", -1, "<--insert text-->"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := i.Insert(tt.args.file, tt.args.line, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("InsertService.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkInsertServiceInsertBegining(t *testing.B) {
	i := &InsertService{}
	fm := io.NewFileManager()
	i.SetManager(fm)
	fm.Save([]string{"Line 1", "Line 2", "Line 3"}, "benchmarktest.txt")
	t.ResetTimer()
	for index := 0; index < t.N; index++ {
		i.Insert("benchmarktest.txt", 1, "<-- insert line -->")
	}
	t.StopTimer()
	os.Remove("benchmarktest.txt")
}

func BenchmarkInsertServiceInsertMiddle(t *testing.B) {
	i := &InsertService{}
	fm := io.NewFileManager()
	i.SetManager(fm)
	fm.Save([]string{"Line 1", "Line 2", "Line 3"}, "benchmarktest.txt")
	t.ResetTimer()
	for index := 0; index < t.N; index++ {
		i.Insert("benchmarktest.txt", 2, "<-- insert line -->")
	}
	t.StopTimer()
	os.Remove("benchmarktest.txt")
}

func BenchmarkInsertServiceInsertEnd(t *testing.B) {
	i := &InsertService{}
	fm := io.NewFileManager()
	i.SetManager(fm)
	fm.Save([]string{"Line 1", "Line 2", "Line 3"}, "benchmarktest.txt")
	t.ResetTimer()
	for index := 0; index < t.N; index++ {
		i.Insert("benchmarktest.txt", -1, "<-- insert line -->")
	}
	t.StopTimer()
	os.Remove("benchmarktest.txt")
}

type fakeFileManager struct {
}

func newFileManager() io.Manager {
	return &fakeFileManager{}
}

func (f *fakeFileManager) Read(file string) ([]string, error) {
	switch file {
	case "file":
		return []string{"line 1", "line 2"}, nil
	case "file3":
		return []string{"line 1", "line 2"}, nil
	case "file2":
		return nil, errors.Wrap(errors.New("Not such file"), "testing")
	}
	return []string{}, nil
}

func (f *fakeFileManager) Save(content []string, file string) error {
	switch file {
	case "file3":
		return errors.Wrap(errors.New("Cannot access file"), "testing")
	}
	return nil
}
