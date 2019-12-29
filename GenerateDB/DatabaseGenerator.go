package main

import (
	"fmt"

	config "github.com/Samuael/Projects/Inovide/DB"
	"github.com/gchaincl/dotsql"
)

func main() {
	db, err := config.InitializPostgres()

	if err != nil {
		panic(err)
	}
	dot, err := dotsql.LoadFromFile("../sql/generateDatabase.sql")
	if err != nil {
		panic(err)
	}
	_, err = dot.Exec(db, "create-database-inoide")
	if err != nil {
		panic(err)
	}
	_, err = dot.Exec(db, "create-users-table")
	if err != nil {
		panic(err)
	}

	fmt.Println("Sucssusfull Created Database Inovide and Table User")

}