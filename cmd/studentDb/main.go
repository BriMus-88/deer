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
	var studentsList []Student
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

//	studentDb.DisplayIntro(dbConnection)

	studentDb.MainMenuDisplay(dbConnection)






	// rows, err := dbConnection.Query(`
	// 	SELECT id, name, course, age, city
	// 	FROM students
	// `)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for rows.Next() {
	// 	var temp Student // A temporary struct for the current row
	// 	err = rows.Scan(&temp.ID, &temp.Name, &temp.Course, &temp.Age, &temp.City)

	// 	// Add this student to our list
	// 	studentsList = append(studentsList, temp)
	// }

	log.Println(studentsList)

}
