package main

import (
	"fmt"
	"go_postgre/config"
	"go_postgre/services"

	_ "go_postgre/controllers"

	_ "github.com/lib/pq"
)

/*
	routing system
*/

func main() {

	fmt.Println("Valdy & Ega")

	_, err := config.DBConnect()

	if err != nil {
		fmt.Println(err)
	}

	// Controllers
	services.InputData()
}
