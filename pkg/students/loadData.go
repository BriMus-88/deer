package students

import (
	"database/sql"

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

var Db *sql.DB

func LoadData() ([]Student, error) {
	var studentsList []Student 
	var err error // Declare err first

	connStr := `host=localhost port=5432 user=postgres password=2811 dbname=studentGolangTraining sslmode=disable`

	Db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := Db.Query(`
		SELECT id, name, course, age, city
		FROM students
	`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var tempS Student // A temporary struct for the current row
		err = rows.Scan(&tempS.ID, &tempS.Name, &tempS.Course, &tempS.Age, &tempS.City)
		if err != nil {
			return nil, err // Return the error if scanning fails
		}
		// Add this student to our list
		studentsList = append(studentsList, tempS)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return studentsList, nil

}


func CountStudents() (int, error) {

	var count int
	var err error

	err = Db.QueryRow(`SELECT COUNT(*) FROM students`).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count, nil	
}