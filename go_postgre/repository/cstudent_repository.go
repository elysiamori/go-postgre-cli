package repository

import (
	"database/sql"
	"fmt"
	"go_postgre/model"
	"strings"
)

type cstudentRepository struct {
	db *sql.DB
}

func NewCStudentRepository(db *sql.DB) *cstudentRepository {
	return &cstudentRepository{db}
}

// receiver (p *pointer)

// CRUD {CREATE}
func (r *cstudentRepository) Add(cstudent *model.CStudent) error {

	// query instert
	query := `INSERT INTO "cstudent"("full_name", "email","age","programming") 
		VALUES($1, $2, $3, $4)`

	queryStat, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer queryStat.Close()

	_, err = queryStat.Exec(cstudent.FullName, cstudent.Email, cstudent.Age, cstudent.Programming)

	if err != nil {
		return err
	}

	fmt.Printf("New Collage Student : %+v\n ", cstudent)

	return nil
}

// CRUD {UPDATE}
func (r *cstudentRepository) Update(id string, cstudent *model.CStudent) error {

	query := `UPDATE "cstudent" SET "full_name"=$1, "email"=$2,"age"=$3,"programming"=$4 
			WHERE "id"=$5`

	queryStat, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer queryStat.Close()

	_, err = queryStat.Exec(cstudent.FullName, cstudent.Email, cstudent.Age, cstudent.Programming, id)

	if err != nil {
		return err
	}

	fmt.Printf("Updated : %+v\n", cstudent)

	return nil
}

// CRUD {DELETE}
func (r *cstudentRepository) Delete(id string) error {

	query := `DELETE FROM "cstudent" WHERE "id" = $1`

	queryStat, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer queryStat.Close()

	_, err = queryStat.Exec(id)

	if err != nil {
		return err
	}

	fmt.Printf("Deleted Collage Student : %+v\n", id)

	return nil
}

// CRUD {READ:ID}
func (r *cstudentRepository) FindById(id string) (*model.CStudent, error) {

	query := `SELECT * FROM "cstudent" WHERE "id"=$1`

	var cstudent model.CStudent

	queryStat, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer queryStat.Close()

	err = queryStat.QueryRow(id).Scan(&cstudent.ID, &cstudent.FullName, &cstudent.Email, &cstudent.Age, &cstudent.Programming)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Collage Student by ID : %+v\n", cstudent)

	return &cstudent, nil
}

// CRUD {READ ALL}
func (r *cstudentRepository) FindAll() (model.CStudents, error) {

	query := `SELECT * FROM "cstudent"`

	var cstudents model.CStudents

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var cstudent model.CStudent

		err = rows.Scan(&cstudent.ID, &cstudent.FullName, &cstudent.Email, &cstudent.Age, &cstudent.Programming)

		if err != nil {
			return nil, err
		}

		cstudents = append(cstudents, cstudent)
	}

	for _, value := range cstudents {
		fmt.Println(strings.Repeat("=", 20))
		fmt.Println("ID          : ", value.ID)
		fmt.Println("Fullname    : ", value.FullName)
		fmt.Println("Email       : ", value.Email)
		fmt.Println("Age         : ", value.Age)
		fmt.Println("Programming : ", value.Programming)
		fmt.Println(strings.Repeat("=", 20))
	}

	return cstudents, nil
}
