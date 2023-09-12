package services

import (
	"fmt"
	"go_postgre/config"
	"go_postgre/controllers"
	"go_postgre/model"
	"go_postgre/repository"
)

// input data

func InputData() {
	var x int
	db, err := config.DBConnect()

	if err != nil {
		fmt.Println(err)
	}

	var (
		a, b, c string
		e       int

		rid string // Read By Id
	)

	vanica := model.NewCStudent()
Loop:
	fmt.Println("Golang PostgreSQL")
	fmt.Println("1.CREATE DATA")
	fmt.Println("2.READ ALL DATA")
	fmt.Println("3.READ BY ID DATA")
	fmt.Println("4.UPDATE DATA")
	fmt.Println("5.DELETE DATA")
	fmt.Println("6.EXIT")
	fmt.Print("Option: ")
	fmt.Scanln(&x)

	switch x {
	case 1:
		fmt.Print("Fullname : ")
		fmt.Scanln(&a)
		fmt.Print("Email : ")
		fmt.Scanln(&b)
		fmt.Print("Age : ")
		fmt.Scanln(&e)
		fmt.Print("Programming : ")
		fmt.Scanln(&c)
		fmt.Println("================")

		vanica.FullName = a
		vanica.Email = b
		vanica.Age = e
		vanica.Programming = c

		cstudentRepository := repository.NewCStudentRepository(db)

		err = controllers.AddCStudent(vanica, cstudentRepository)

		if err != nil {
			fmt.Println(err)
		}
		goto Loop

	case 2:
		cstudentRepository := repository.NewCStudentRepository(db)
		_, err = controllers.FindallCStudent(cstudentRepository)

		if err != nil {
			fmt.Println(err)
		}
		goto Loop

	case 3:
		cstudentRepository := repository.NewCStudentRepository(db)
		fmt.Print("Masukan ID: ")
		fmt.Scanln(&rid)
		_, err = controllers.FindbyidCStudent(rid, cstudentRepository)

		if err != nil {
			fmt.Println(err)
		}
		goto Loop

	case 4:
		cstudentRepository := repository.NewCStudentRepository(db)
		fmt.Print("Masukan ID: ")
		fmt.Scanln(&rid)
		vanica, err := controllers.FindbyidCStudent(rid, cstudentRepository)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Fullname : ")
		fmt.Scanln(&a)
		fmt.Print("Email : ")
		fmt.Scanln(&b)
		fmt.Print("Age : ")
		fmt.Scanln(&e)
		fmt.Print("Programming : ")
		fmt.Scanln(&c)
		fmt.Println("================")

		vanica.FullName = a
		vanica.Email = b
		vanica.Age = e
		vanica.Programming = c

		err = controllers.UpdateCStudent(vanica, cstudentRepository)

		if err != nil {
			fmt.Println(err)
		}
		goto Loop

	case 5:
		cstudentRepository := repository.NewCStudentRepository(db)
		fmt.Print("Masukan ID: ")
		fmt.Scanln(&rid)
		err = controllers.DeleteCStudent(rid, cstudentRepository)

		if err != nil {
			fmt.Println(err)
		}
		goto Loop

	default:
		fmt.Println("Thanks")
	}

}
