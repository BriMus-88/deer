package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/briMus-88/deer/pkg/studentDb"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Student struct {
	ID     int
	Name   string
	Course string
	Age    int
	City   string
}


func main() {
	var dbConnection *sql.DB
	connStr := `host=localhost port=5432 user=postgres password=2811 dbname=studentGolangTraining sslmode=disable`

	dbConnection, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	count := studentDb.CountStudents(dbConnection)
	fmt.Println("Number of students:", count)

	studentDb.DisplayIntro()

	studentDb.MainMenuDisplay(dbConnection)

	defer dbConnection.Close()
}
