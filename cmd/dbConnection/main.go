package main

import (
	"database/sql"
	"fmt"
	"log"
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
	connStr := `host=localhost port=5432 user=postgres password=2811 dbname=studentGolangTraining sslmode=disable`

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`
		SELECT id, name, course, age, city
		FROM students
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s Student
		err := rows.Scan(&s.ID, &s.Name, &s.Course, &s.Age, &s.City)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
