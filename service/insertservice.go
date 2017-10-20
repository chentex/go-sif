package service

//InsertServiceInterface defines the methods that any inserter need to follow
type InsertServiceInterface interface {
	Insert(file string, line int, text string) error
}

//InsertService Implementing a file based insert service
type InsertService struct {
}

//NewInsertService return a new instance of InsertService
func NewInsertService() InsertServiceInterface {
	return &InsertService{}
}

//Insert text into file in given line
func (i *InsertService) Insert(file string, line int, text string) error {
	return nil
}
