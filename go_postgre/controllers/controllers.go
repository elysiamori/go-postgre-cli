package controllers

import (
	_ "database/sql"
	"go_postgre/model"
	"go_postgre/repository"
)

// controllers
// ADD {CRUD:CREATE}
func AddCStudent(p *model.CStudent, repo repository.CStudentRepository) error {
	err := repo.Add(p)

	if err != nil {
		return err
	}

	return nil
}

// UPDATE {CRUD:UPDATE}
func UpdateCStudent(p *model.CStudent, repo repository.CStudentRepository) error {
	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}

	return nil
}

// DELETE {CRUD:DELETE}
func DeleteCStudent(id string, repo repository.CStudentRepository) error {
	err := repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// READ BY ID {CRUD:READ BY ID}
func FindbyidCStudent(id string, repo repository.CStudentRepository) (*model.CStudent, error) {
	cstudent, err := repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return cstudent, nil
}

// READ ALL {CRUD:READ ALL}
func FindallCStudent(repo repository.CStudentRepository) (model.CStudents, error) {
	cstudents, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return cstudents, nil
}
