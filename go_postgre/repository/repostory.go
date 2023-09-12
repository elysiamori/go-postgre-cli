package repository

import (
	"go_postgre/model"
)

// abstraction cstudent repository
// untuk unit test atau memakai map
type CStudentRepository interface {
	Add(*model.CStudent) error
	Update(string, *model.CStudent) error
	Delete(string) error
	FindById(string) (*model.CStudent, error)
	FindAll() (model.CStudents, error)
}
